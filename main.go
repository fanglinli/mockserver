package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fanglinli/mockserver/mocker"
)

func main() {
	// Parse listen port and API response specification files location.
	listen := flag.String("listen", ":8080", "Port to listen on.")
	filesLocation := flag.String("folder", ".", "Folder location of response specification file.")
	flag.Parse()

	files, err := ioutil.ReadDir(*filesLocation)
	if err != nil {
		log.Fatal(err)
	}

	// Register response handlers for each API response specification file.
	fileNames := make([]string, len(files))
	for i, file := range files {
		mocker.HandleRequestForFile(fmt.Sprintf("%s/%s", *filesLocation, file.Name()), file.Name())
		fileNames[i] = file.Name()
	}

	// Start the mock endpoint(s).
	fmt.Println(fmt.Sprintf("Listening on port: %s", (*listen)[1:]))
	for _, fileName := range fileNames {
		fmt.Println(fmt.Sprintf("Serving endpoint for file: %s", fileName))
	}
	http.ListenAndServe(*listen, nil)
}
