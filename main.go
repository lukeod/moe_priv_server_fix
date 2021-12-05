package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func get_private_list_file() string {
	content, err := ioutil.ReadFile("server_list.json")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Responding to /private_list")
	return string(content)
}

func private_list(w http.ResponseWriter, req *http.Request) {
	var response = get_private_list_file()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, response)
}

func main() {
	http.HandleFunc("/private_list", private_list)
	http.ListenAndServe(":80", nil)
}
