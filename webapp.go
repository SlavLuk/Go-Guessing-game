package main

import (
	
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w, "Guessing game")//print out to the page
	
}

func main(){


	http.HandleFunc("/", handler)//handle any request 
    http.ListenAndServe(":8080", nil)//listen to port 8080 in infinite loop
}