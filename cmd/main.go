package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"Hotel_reservation/internal/api"
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
	http.HandleFunc("/signup", api.Signup)
	http.HandleFunc("/login",api.LoginHandler)
	
	fmt.Println("Server started on :8080")
	http.ListenAndServe(": 8080" , nil)
	
}