package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	configD := os.Getenv("CONFIGD")

	projectPath := flag.String("path", "./", "Path to project root.")
	configKey := flag.String("key", configD, "Config'D API Key")

	isNode := flag.Bool("node", false, "Specify if launcher should use NodeJS recipe.")
	isGo := flag.Bool("go", false, "Specify if launcher should use Go recipe.")

	flag.Parse()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("TOKEN", *configKey)
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		err := errors.New("Error, please verify your API key set. ")
		log.Fatal(err)
	}

	body := ReadBody(res)
	err = os.Chdir(*projectPath)

	if err != nil {
		log.Fatal(err)
	}

	if *isNode {
		ExportForNode(body)
	}

	if *isGo {
		ExportForGo(body)
	}

}
