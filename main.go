package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(executablePath)
	fileserver := http.FileServer(http.Dir(dir))

	http.Handle("/", fileserver)

	log.Println(" listen on port 9088")
	if err := http.ListenAndServe(":9088", nil); err != nil {
		log.Fatal(err)
	}
}
