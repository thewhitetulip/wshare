package main

import  (
	"flag"
	"net/http"
	"os"
	// "fmt"
	)

var filename *string
var numberOfTimes *int
var downloaded int

func main() {
	downloaded = 0
	filename = flag.String("f", "invalid", "File to be shared")
	numberOfTimes = flag.Int("t", 1, "Number of times to be shared")
	flag.Parse()
	http.HandleFunc("/share/", ServeFileHandler)
	http.ListenAndServe(":8080", nil)
}

func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	if downloaded <= numberofTimes {
		http.ServeFile(w, r, "./" + *filename)
		downloaded += 1
	} else {
		os.Exit(0)
	}
}