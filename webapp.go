package main

import (
	
	"net/http"
	"html/template"
)

type Page struct{

	Message string
}

func handler(w http.ResponseWriter, r *http.Request){

 var p= &Page{Message:"Guessing game"}//assign title
 
 t,_ := template.ParseFiles("home.html")//parse html page

	t.Execute(w,p)
	
}

func guessHandler(w http.ResponseWriter, r *http.Request){

	var p = &Page{Message:"Guess a number between 1 and 20"}
	t,_ := template.ParseFiles("guess.tmpl")//parse html page
	t.Execute(w,p)
}

func main(){


	http.HandleFunc("/", handler)//handle any request 
	http.HandleFunc("/guess",guessHandler)//handle /guess request
    http.ListenAndServe(":8080", nil)//listen to port 8080 in infinite loop
}