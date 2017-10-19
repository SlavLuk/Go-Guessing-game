package main

import (
	
	"net/http"
	"html/template"
	"time"
	"math/rand"
	"log"
	"strconv"

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

	var cookievalue string

	randNum := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(20)+1//generate random 1 <= number <= 20 
	
	randStr := strconv.Itoa(randNum)//convert random int to string

	cookie,err:= r.Cookie("target")//get cookie named target

	if err !=nil{

		log.Printf("Error occured %s",err)//log error
	
	}else{

		cookievalue = cookie.Value
	
	}


	 if cookievalue == ""{//check if cookie is set

		cookieNew := http.Cookie{Name:"target",Value:randStr}//create cookie target to rand number
	
		http.SetCookie(w,&cookieNew)//set cookie
	}
    
	p := &Page{Message:"Guess a number between 1 and 20"}//create struct Page

	t,_ := template.ParseFiles("guess.tmpl")//parse html page
	t.Execute(w,p)
}

func main(){


	http.HandleFunc("/", handler)//handle any request 
	http.HandleFunc("/guess",guessHandler)//handle /guess request
    http.ListenAndServe(":8080", nil)//listen to port 8080 in infinite loop
}