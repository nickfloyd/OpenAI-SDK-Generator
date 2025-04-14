package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var version string

// Docs [here](https://pkg.go.dev/flag) recommends this approach.
func init() {
	flag.StringVar(&version, "version", "", "Assumes OpenAI API is versioned by the version of the OpenAPI spec.")
}

func main() {
	if err := realMain(); err != nil {
		log.Fatal(err)
	}
	log.Println("execution finished successfully")
}

func realMain() error {
	flag.Parse()

	downloadedFile := "schemas/openai-openapi.yaml"
	fileName := "openapi.yaml"

	out, err := os.Create(downloadedFile)
	if err != nil {
		return err
	}

	defer out.Close()

	logMsg := "Downloading latest schema from descriptions directory"
	url := "https://raw.githubusercontent.com/openai/openai-openapi/refs/heads/master/" + fileName

	log.Printf(logMsg)
	log.Printf("OpenAI OpenAPI source url: %s", url)

	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		err := os.Remove(downloadedFile)
		if err != nil {
			return fmt.Errorf("Error deleting file: %s. Error: %s", downloadedFile, err)
		}
		return fmt.Errorf("Received 404 Not Found for url: %s", url)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
