package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type doggo struct {
	Name  string
	Breed string
	Age   int
}

type catto struct {
	Name  string `json:"catName"`
	Breed string `json:"catType"`
	Age   int    `json:"catAge"`
}

func main() {
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/unmarshal", unmarshal)
	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	http.HandleFunc("/get", get)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func marshal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	d := doggo{"Marshall", "mutt", 5}
	json, err := json.Marshal(d)
	if err != nil {
		log.Println(err)
	}
	w.Write(json)
}

func encode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	d := doggo{"Encoder", "mutt", 5}
	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		log.Println(err)
	}
}

func unmarshal(w http.ResponseWriter, r *http.Request) {
	var d doggo
	j := `{"Name":"Unmarshal","Breed":"mutt","Age":5}`
	err := json.Unmarshal([]byte(j), &d)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(d)
	w.WriteHeader(http.StatusOK)
}

func decode(w http.ResponseWriter, r *http.Request) {
	var c catto
	j := `{"catName":"Unmarshal Catto","catType":"Persian","catAge":6}`
	err := json.NewDecoder(strings.NewReader(j)).Decode(&c)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(c)
	w.WriteHeader(http.StatusOK)
}

func get(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080/marshal")
	if err != nil {
		fmt.Println(err)
	}

	j := make([]byte, resp.ContentLength)

	defer resp.Body.Close()

	_, err = resp.Body.Read(j)
	if err != io.EOF {
		fmt.Println(err)
	}

	fmt.Println(string(j))
}
