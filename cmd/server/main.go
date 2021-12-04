//Server for testing purposes
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	ServerPort = ":8080"
	SourceDir  = "/home/paulidealiste/gop/SnapCooter/web"
)

func wd() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func main() {
	fserver := http.FileServer(http.Dir(SourceDir))

	fmt.Printf("Server running from %s", wd())

	if err := http.ListenAndServe(ServerPort, fserver); err != nil {
		log.Fatal(err)
	}
}
