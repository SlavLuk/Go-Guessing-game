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
	Guess int
	Link string
	MessageWin string

}
func randomGen()string{

	randNum := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(20)+1//generate random 1 <= number <= 20 
	
	randStr := strconv.Itoa(randNum)//convert random int to string

	return randStr

}
func handler(w http.ResponseWriter, r *http.Request){

 p := &Page{Message:"Guessing game"}//assign title
 
 t,_ := template.ParseFiles("home.html")//parse html page

	t.Execute(w,p)
	
}

func guessHandler(w http.ResponseWriter, r *http.Request){

	var cookievalue string
	var g int
	var randStr string

	

	cookie,err:= r.Cookie("target")//get cookie named 'target'

	if err !=nil{

		log.Printf("Error occured %s",err)//log error
	
	}else{

		cookievalue = cookie.Value
	
	}


	 if len(cookievalue) == 0{//check if cookie is set

		randStr =randomGen()

		cookieNew := http.Cookie{Name:"target",Value:randStr}//create cookie target to rand number
	
		http.SetCookie(w,&cookieNew)//set cookie
	}


		urlVal := r.FormValue("guess")//get guess value from input form


	 if len(urlVal) == 0{//check if guess is set

		 urlVal = "0"//set to 0
	}

	p := &Page{Message:"Guess a number between 1 and 20"}//create struct Page

	g,_ = strconv.Atoi(urlVal)//convert string to int
	cv,_ :=strconv.Atoi(cookievalue)

 	p.Guess = g//assign int to struct var guess

	if cv > g && g!=0{
		
			p.MessageWin = "Your guess was too low"

		}else if cv<g{

			p.MessageWin = "Your guess was too high"

		}else if g == cv && g!=0{

			randStr = randomGen()
			cookieNew := http.Cookie{Name:"target",Value:randStr}//create cookie target to rand number
			
			http.SetCookie(w,&cookieNew)//set cookie

			p.MessageWin = "Congratulations You won !!!"
			p.Link = "New Game"
	
		}

			
	
	t,_ := template.ParseFiles("guess.tmpl")//parse html page
	t.Execute(w,p)
}

func main(){


	http.HandleFunc("/", handler)//handle any request 
	http.HandleFunc("/guess",guessHandler)//handle /guess request
    http.ListenAndServe(":8080", nil)//listen to port 8080 in infinite loop
}