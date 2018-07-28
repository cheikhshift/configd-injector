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
	
	// Set program's default API url with
	// environment variable CONFIGD_URL,
	// if value is present  
	if hostName := os.Getenv("CONFIGD_URL"); hostName != "" {
		apiURL = hostName
	}



	projectPath := flag.String("path", "./", "Path to project root.")
	configKey := flag.String("key", configD, "Config'D API Key")

	isNode := flag.Bool("node", false, "Specify if injector should use NodeJS recipe.")
	isGo := flag.Bool("go", false, "Specify if injector should use Go recipe.")
	isJava := flag.Bool("java", false, "Specify if injector should use Java recipe.")

	flag.Parse()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Set("TOKEN", *configKey)
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		err := errors.New("Error, please verify your API key. ")
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

	if *isJava {
		ExportForJava(body)
	}

	if *isGo {
		ExportForGo(body)
	}

}
