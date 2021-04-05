package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w," Post request successful")
	name := r.FormValue("name")
	addr := r.FormValue("address")

	fmt.Fprint(w, " Name:", name )
	fmt.Fprint(w, " Address:", addr )
}

func main(){
	fmt.Println("Starting server at port 8080")

	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/form",formHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}