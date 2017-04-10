package main

import (
	"encoding/json"
	"net/http"
	"unicode/utf8"
	"log"
	"fmt"
	"github.com/julienschmidt/httprouter"
)

type Output struct {
	Out  string `json:"output"`
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func Rev(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    //fmt.Fprintf(w, "Input:, %s!\n\n", ps.ByName("input"))

	s := ps.ByName("input")
	
	totalLength := len(s)
    buffer := make([]byte, totalLength)
    for i := 0; i < totalLength; {
        r, size := utf8.DecodeRuneInString(s[i:])
        i += size
        utf8.EncodeRune(buffer[totalLength-i:], r)
    }

    out := Output{string(buffer)}

    js, err := json.Marshal(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)


}

func main() {
    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/reverse/:input", Rev)

    log.Fatal(http.ListenAndServe(":8080", router))
}
