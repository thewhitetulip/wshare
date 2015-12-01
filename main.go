package main

import (
	"flag"
	"fmt"
	"strings"
	"net/http"
	"net"
	"os"
)

var filename *string
var numberOfTimes *int
var downloaded int

func main() {
	downloaded = 0
	addresses, err := net.InterfaceAddrs()
	checkErr(err)
	
	ip := strings.Split(addresses[0].String(), "/")[0]
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

func checkErr(err error){
	if err!= nil {
		fmt.Println("Error: ", err)
	}
}