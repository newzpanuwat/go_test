package main

import (
		"encoding/json"
		"net/http"
		)

type contact struct {
	Name string `json:"name"`
	Tel string	`json:"tel"`
}

func (c *contact) String() string {
	return c.Name + ", " + c.Tel
}

func main() {
	http.HandleFunc("/", get)
	http.HandleFunc("/savecontact", post)
	http.Handle("/contact", ch)
	http.ListenAndServe(":8080", nil)

}

//GET
func get(w http.ResponseWriter, r *http.Request) {
	c := contact { Name: "gopher", Tel: "0123456789" }
	json.NewEncoder(w).Encode(c)
}

//POST
func post(w http.ResponseWriter, r *http.Request){
	var c contact
	json.NewDecoder(r.Body).Decode(&c)

	fmt.Println(c)
}

