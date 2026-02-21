package main

import (
	"Hotel_reservation/internal/user"
	"Hotel_reservation/config"
	"Hotel_reservation/db/postgres"
	"Hotel_reservation/internal/api"
	"fmt"
	"log"
	"net/http"
	"os"
)
var (
	ErrorLogger	 *log.Logger
	WarnLogger   *log.Logger
	InfoLogger   *log.Logger
)

func init(){
	file, err := os.OpenFile("log.txt", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil{
		log.Fatal("Failed to open log file" , err)
	}

	flags := log.Ldate | log.Ltime | log.Lshortfile
	log.SetFlags(flags)
	log.SetOutput(file)

	ErrorLogger = log.New(file, "Error :", flags)
	WarnLogger  = log.New(file, "Warn :" , flags)
	InfoLogger  = log.New(file, "info :" , flags)
}




func main(){

	config.LoadDB()
	dsn := config.PostgresDSN()
	db, err := postgres.InitPostgres(dsn)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer db.Close()

	userstore := postgres.NewPostgresDB(db)
	userservice := user.NewService(userstore)
	userhandler := api.NewHandler(userservice)

	http.HandleFunc("/signup", userhandler.Signup)
	http.HandleFunc("/login",userhandler.LoginHandler)
	
	fmt.Println("Server started on :8080")
	http.ListenAndServe(": 8080" , nil)
	
}