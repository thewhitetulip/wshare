package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var filename *string
var numberOfTimes *int
var downloaded int

func main() {
	downloaded = 0
	ip := "127.0.0.1"
	//TODO -> Get the system IP
	filename = flag.String("f", "invalid", "File name to be shared")
	numberOfTimes = flag.Int("t", 1, "Number of times file to be shared")

	flag.Parse()
	http.HandleFunc("/share/", ServeFileHandler)

	if *filename == "invalid" {
		fmt.Println("Invalid file, exiting")
		flag.PrintDefaults()
		os.Exit(0)
	}
	//TODO share a simple link like :8080/r/1 and redirect that to /share/classical.pdf
	fmt.Printf("download link: http://%s:8080/share/%s\n", ip, *filename)
	fmt.Println("Running server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/share/"):]
	downloaded += 1
	http.ServeFile(w, r, "./"+file)
	if downloaded >= *numberOfTimes {
		fmt.Println("Downloads done, exiting")
		os.Exit(0)
	}
}
