package main

import (
	"archive/zip"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var filename *string
var numberOfTimes *int
var downloaded int
var routes map[string]string

func main() {
	downloaded = 0
	routes = make(map[string]string)
    ip := GetOutboundIP().String() 
	filename = flag.String("f", "invalid", "File name to be shared")
	numberOfTimes = flag.Int("t", 1, "Number of times file to be shared")

	flag.Parse()
	http.HandleFunc("/r/", ServeRouterHandler)
	http.HandleFunc("/share/", ServeFileHandler)

	if *filename == "invalid" {
		log.Println("Invalid file, exiting")
		flag.PrintDefaults()
		os.Exit(0)
	}

	file, err := os.Open(*filename)
	checkErr(err)
	fileStat, err := file.Stat()
	checkErr(err)

	if fileStat.IsDir() {
		zipname := *filename + ".zip"
		zipit(*filename, zipname)
		routes["1"] = zipname
	} else {
		routes["1"] = *filename
	}

	log.Printf("download link: http://%s:8080/r/1\n", ip)
	log.Println("Running server on port 8080")
	http.ListenAndServe(":8080", nil)
}

//ServeRouterHandler Used to serve the simpe URL router
//the parameter passed is an integer which is looked up on a map to fetch the name
func ServeRouterHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/r/"):]
	http.Redirect(w, r, "/share/"+routes[path], http.StatusFound)
}

//ServeFileHandler Used to serve the actual file
func ServeFileHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path[len("/share/"):]
	downloaded += 1
	log.Println("Downloaded", downloaded, " times")
	http.ServeFile(w, r, file)
	if downloaded >= *numberOfTimes {
		log.Println("Downloads done, exiting")
		//os.Exit(0)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}

//zipit creates a "target".zip file from the source directory
func zipit(source, target string) {
	zipfile, err := os.Create(target)
	checkErr(err)
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	checkErr(err)
	var baseDir string

	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		header, err := zip.FileInfoHeader(info)
		checkErr(err)

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		checkErr(err)

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		defer file.Close()

		_, err = io.Copy(writer, file)

		if err != nil {
			return nil
		}

		return err

	})
}

// extracts the IP address in more efficient way
func GetOutboundIP() net.IP {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close() 
    localAddr := conn.LocalAddr().(*net.UDPAddr)
    return localAddr.IP
}
