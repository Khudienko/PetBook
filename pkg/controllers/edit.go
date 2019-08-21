package controllers

import (
	"fmt"
	"github.com/dpgolang/PetBook/pkg/models"
	"github.com/dpgolang/PetBook/pkg/view"
	"github.com/gorilla/context"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Editstr struct {
	Name string
	Email string
	Password string
	PetName string
	Age string
	PetType string
	Breed string
	Description string
	Weight string
	Gender string
}

type UserChar struct {
	Name string
	Email string
	Password string
}

func (c *Controller) EditHandler(w http.ResponseWriter, r *http.Request) {
	id := context.Get(r, "id").(int)
	user,_:=c.UserStore.GetUser(id)
	var usChar UserChar
	usChar.Name=user.Firstname
	usChar.Email=user.Email
	usChar.Password=user.Password

	pet, _ := c.UserStore.GetPet(id)
	var edit Editstr
	edit.Name=user.Firstname
	edit.Email=user.Email
	edit.Password=user.Password
	edit.PetName=pet.Name
	edit.Age=pet.Age
	edit.PetType=pet.PetType
	edit.Breed=pet.Breed
	edit.Description=pet.Description
	edit.Weight=pet.Weight
	edit.Gender=pet.Gender

	view.GenerateHTML(w,"Settings","navbar")
	view.GenerateHTML(w,edit,"edit")
	//view.GenerateHTML(w,nil,"upload")
	view.GenerateHTML(w,nil,"footer")
}
func (c *Controller) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	id := context.Get(r, "id").(int)
	if err != nil {
		log.Println(err)
	}
	pet:= &models.Pet{}
	pet.ID=id
	pet.Name=r.FormValue("name")
	pet.PetType=     r.FormValue("animal_type")
	pet.Breed=       r.FormValue("breed")
	pet.Age=      r.FormValue("age")
	pet.Weight =   r.FormValue("weight")
	pet.Gender=      r.FormValue("gender")
	pet.Description= r.FormValue("description")
	c.PetStore.UpdatePet(pet)

	http.Redirect(w, r, "/mypage", 301)
}

func (c *Controller) GetImgHandler(w http.ResponseWriter, r *http.Request) {
	adr:=GetImg()
	view.GenerateHTML(w,adr,"getimg")
}
//
func GetImg() []string {
	var files []string

	root := "./web/static/"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".jpg" && filepath.Ext(path) != ".png" {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}

	//fmt.Println(files[1:])
	for i,file:=range files{
		file=files[i]
		v:=strings.Replace(file,"\\","/",100)
		files[i]=v
	}
	for i,file:=range files{
		file=files[i]
		v:=strings.Replace(file,"web","..",100)
		files[i]=v
	}
	//return files[1:]
	return files
}


//func GetImg()[]os.FileInfo  {
//	files, err := ioutil.ReadDir("./web/storage/")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, file := range files {
//
//		fmt.Println(file.IsDir())
//		fmt.Println(file.Sys())
//
//	}
//
//	return files
//}