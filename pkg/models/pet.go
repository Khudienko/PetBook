package models

import (
	//"github.com/dpgolang/PetBook/pkg/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type Pet struct {
	ID          int    `db:"user_id"`
	Name        string `db:"name"`
	PetType     string `db:"animal_type"`
	Breed       string `db:"breed"`
	Age         string `db:"age"`
	Weight      string `db:"weight"`
	Gender      string `db:"gender"`
	Description string `db:"description"`
}

type PetStore struct {
	DB *sqlx.DB
}

type PetStorer interface {
	RegisterPet(pet *Pet) error
	UpdatePet(pet *Pet)
	GetPetEnums()[]string
}

// TODO: rewrite to update into
func (p *PetStore) RegisterPet(pet *Pet) error {
	_, err := p.DB.Exec("insert into pets (user_id, name, animal_type, breed, age, weight, gender, description) values ($1, $2, $3, $4, $5, $6, $7, $8)",
		pet.ID, pet.Name, pet.PetType, pet.Breed, pet.Age, pet.Weight, pet.Gender, pet.Description)
	if err != nil {
		return fmt.Errorf("cannot affect rows in pets in db: %v", err)
	}
	return nil
}

func (p *PetStore) UpdatePet(pet *Pet)  {
	_, err := p.DB.Exec("update pets set name=$1, age=$2,animal_type=$3, breed =$4,weight=$5,gender=$6,description=$7 where user_id = $8",
		pet.Name, pet.Age, pet.PetType, pet.Breed,pet.Weight,pet.Gender,pet.Description, pet.ID)
	//_, err := p.DB.Exec("update pets set name=$1 where user_id = $2", pet.Name, pet.ID)
	if err != nil {
		log.Println(err)
	}
}
func (p *PetStore) GetPetEnums() []string {
	var petType []string
	var ptype string
	rows,err:=p.DB.Queryx("select unnest(enum_range(NULL::kind_of_animal))::text")
	if err != nil {
		fmt.Println("Error in getting enums")
	}
	for rows.Next(){
		err = rows.Scan(&ptype)
		if err != nil{
			logFatal(err)
		}
		petType=append(petType,ptype)
	}
	fmt.Println(petType)
	return petType
}
