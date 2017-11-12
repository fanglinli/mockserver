package mocker

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HandleRequestForFile registers response handlers for each file specification.
// The pattern for each handler is the base name of the given file.
func HandleRequestForFile(filePath, fileName string) {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc(
		fmt.Sprintf("/%s", baseName(fileName)),
		createHandler(string(fileContents)),
	)
}

// baseName returns the name of a file without the extension.
func baseName(fileName string) string {
	for i := 0; i < len(fileName)-1 && fileName[i] != '/'; i++ {
		if fileName[i] == '.' {
			return fileName[:i]
		}
	}
	return fileName
}

// createHandler creates a handler function from the given format string.
func createHandler(fileContents string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, fmt.Sprintf(fileContents, returnRandom(fileContents)...))
	}
}
