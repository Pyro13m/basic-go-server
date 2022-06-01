package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) { //any HTTP route/api has a res(w),req(r)
	if r.URL.Path != "/hello" { //if any anomaly in the path
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" { //default method of GET
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() Error %v", err)
		return
	}
	fmt.Fprintf(w, "POST Request Successful")
	name := r.FormValue("name")
	age := r.FormValue("age")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Age = %v\n", age)

}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server Started at Port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
