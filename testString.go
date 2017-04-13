package main

import (
	"encoding/json"
	"net/http"
	"unicode/utf8"
	//"log"
	//"fmt"
	"io/ioutil"

	"github.com/gorilla/mux"
)

type Input struct {
	Input string
}
type Output struct {
	Output string
}

func reverse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8") //set for unicode

	var in Input

	b, err := ioutil.ReadAll(r.Body) //read the body
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &in) //store data from body into struct

	// call reverse string here
	out := Output{reverseString(in.Input)}

	j, _ := json.MarshalIndent(out, "", "\t") // pretty print json data back
	w.Write([]byte("\n"))
	w.Write(j)
	w.Write([]byte("\n\n"))

}

func reverseString(s string) string {
	totalLength := len(s)
	buffer := make([]byte, totalLength)
	for i := 0; i < totalLength; {
		r, size := utf8.DecodeRuneInString(s[i:]) //each character has runic bytes that contains unicode
		i += size
		utf8.EncodeRune(buffer[totalLength-i:], r) //reverse here and encode the runic unicode byte back
	}
	return string(buffer)
}

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/echo", echo).Methods("POST")
	r.HandleFunc("/reverse", reverse).Methods("POST")
	http.ListenAndServe(":8080", nil)
}

func echo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8") //set for unicode

	var in Input

	b, err := ioutil.ReadAll(r.Body) //read the body
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &in) //store data from body into struct

	out := Output{in.Input}

	j, _ := json.MarshalIndent(out, "", "\t") // pretty print json data back
	w.Write([]byte("\n"))
	w.Write(j)
	w.Write([]byte("\n\n"))

}
