package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"unicode"
	"unicode/utf8"

	"github.com/gorilla/mux"
)

type Input struct {
	Input string
}
type Output struct {
	Output string
}

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/echo", echo).Methods("POST")
	r.HandleFunc("/reverse", reverse).Methods("POST")
	http.ListenAndServe(":80", nil)
}

func reverse(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json; charset=utf-8") //set for unicode

	var in Input

	b, err := ioutil.ReadAll(r.Body) //read the body
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &in) //store data from body into struct

	// call reverse string here
	out := Output{reversePreservingCombiningCharacters(in.Input)}

	j, _ := json.MarshalIndent(out, "", "\t") // pretty print json data back
	w.Write([]byte("\n"))
	w.Write(j)
	w.Write([]byte("\n\n"))

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

func reversePreservingCombiningCharacters(s string) string {
	if s == "" {
		return ""
	}
	p := []rune(s)
	r := make([]rune, len(p))
	start := len(r)
	for i := 0; i < len(p); {
		// quietly skip invalid UTF-8
		if p[i] == utf8.RuneError {
			i++
			continue
		}
		j := i + 1
		for j < len(p) && (unicode.Is(unicode.Mn, p[j]) ||
			unicode.Is(unicode.Me, p[j]) || unicode.Is(unicode.Mc, p[j])) {
			j++
		}
		for k := j - 1; k >= i; k-- {
			start--
			r[start] = p[k]
		}
		i = j
	}
	return (string(r[start:]))
}
