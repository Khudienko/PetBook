package main

import (
	"github.com/dpgolang/PetBook/gomigrations"
	"github.com/dpgolang/PetBook/pkg/authentication"
	"github.com/dpgolang/PetBook/pkg/controllers"
	"github.com/dpgolang/PetBook/pkg/driver"
	"github.com/dpgolang/PetBook/pkg/logger"
	_ "github.com/dpgolang/PetBook/pkg/logger"
	"github.com/dpgolang/PetBook/pkg/models"
	"github.com/dpgolang/PetBook/pkg/models/forum"
	"github.com/dpgolang/PetBook/pkg/models/search"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/urfave/negroni"

	"net/http"
	"os"
)

func main() {

	db := driver.ConnectDB()
	err := gomigrations.Migrate(db)
	if err != nil {
		logger.FatalError(err, "Migration failed.\n")
	}

	router := mux.NewRouter()

	storeUser := models.UserStore{DB: db}
	storePet := models.PetStore{DB: db}
	storeForum := forum.ForumStore{DB: db}
	storeSearch := search.SearchStore{DB:db}
	storeBlog := models.BlogStore{DB: db}

	controller := controllers.Controller{
		PetStore:  &storePet,
		UserStore: &storeUser,
		ForumStore: &storeForum,
		SearchStore: &storeSearch,
		BlogStore: &storeBlog,
	}

	router.HandleFunc("/register", controller.RegisterPostHandler()).Methods("POST")
	router.HandleFunc("/register", controller.RegisterGetHandler()).Methods("GET")

	router.HandleFunc("/login", controller.LoginPostHandler()).Methods("POST")
	router.HandleFunc("/login", controller.LoginGetHandler()).Methods("GET")

	subrouter := router.PathPrefix("/").Subrouter()
	subrouter.Use(mux.MiddlewareFunc(authentication.Content))
	router.HandleFunc("/upload",controllers.UploadFile)
	router.HandleFunc("/petcabinet",controller.PetGetHandler()).Methods("Get")
	router.HandleFunc("/petcabinet",controller.PetPostHandler()).Methods("POST")

	//router.HandleFunc("/mypage",controller.IndexHandler)


	subrouter.HandleFunc("/mypage", controller.MyPageGetHandler()).Methods("GET")
	subrouter.HandleFunc("/petcabinet", controller.PetPostHandler()).Methods("POST")
	subrouter.HandleFunc("/petcabinet", controller.PetGetHandler()).Methods("GET")
	subrouter.HandleFunc("/forum", controller.ViewTopicsHandler()).Methods("GET")
	subrouter.HandleFunc("/forum/submit", controller.NewTopicHandler()).Methods("POST")
	subrouter.HandleFunc("/forum/submit", controller.NewTopicHandler()).Methods("GET")
	subrouter.HandleFunc("/search", controller.ViewSearchHandler()).Methods("GET")
	subrouter.HandleFunc("/search", controller.SearchHandler()).Methods("POST")

	//router.Handle("/", negroni.New(
	//	negroni.HandlerFunc(authentication.ValidateTokenMiddleware),
	//	negroni.Wrap(http.HandlerFunc(controller.MyPageGetHandler())),
	//))
	router.HandleFunc("/",controller.GetBlogHandler)
	router.HandleFunc("/process",controller.CreateBlogHandler)
	http.HandleFunc("/delete",controller.DeleteBlogHandler)







	//router.Handle("/mypage", negroni.New(
	//	negroni.HandlerFunc(authentication.ValidateTokenMiddleware),
	//	negroni.Wrap(http.HandlerFunc(controller.MyPageGetHandler())),
	//)).Methods("GET")
	//
	//router.Handle("/petcabinet", negroni.New(
	//	negroni.HandlerFunc(authentication.ValidateTokenMiddleware),
	//	negroni.Wrap(http.HandlerFunc(controller.PetPostHandler())),
	//)).Methods("POST")
	//
	//router.Handle("/petcabinet", negroni.New(
	//	negroni.HandlerFunc(authentication.ValidateTokenMiddleware),
	//	negroni.Wrap(http.HandlerFunc(controller.PetGetHandler())),
	//)).Methods("GET")
	//
	//router.Handle("/forum", negroni.New(
	//	negroni.HandlerFunc(authentication.ValidateTokenMiddleware),
	//	negroni.Wrap(http.HandlerFunc(controller.ViewTopicsHandler())),
	//)).Methods("GET")
	//
	//router.Handle("/forum/new_topic", negroni.New(
	//	negroni.HandlerFunc(authentication.ValidateTokenMiddleware),
	//	negroni.Wrap(http.HandlerFunc(controller.NewTopicHandler())),
	//))
	//
	//router.Handle("/", negroni.New(
	//	negroni.HandlerFunc(authentication.ValidateTokenMiddleware),
	//	negroni.Wrap(http.HandlerFunc(controller.MyPageGetHandler())),
	//))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./web/static/"))))

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	// Is it proper way to handle ListenAndServe() error?
	if err:= http.ListenAndServe(":8080", loggedRouter); err !=nil {
		logger.FatalError(err, "Error occurred, while trying to listen and serve a server")
	}
}