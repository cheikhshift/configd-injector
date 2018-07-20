package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const NodeTemplate = `// File genereated by Configd.
var _c = %s

module.exports = _c.Config

`

func ExportForNode(config string) {
	CreateDepFolder()

	moduleStr := fmt.Sprintf(NodeTemplate, config)
	module := []byte(moduleStr)
	err := ioutil.WriteFile("node_modules/configd.js", module, 0700)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateDepFolder() {
	if _, err := os.Stat("node_modules"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err = os.Mkdir("node_modules", 0700)
		if err != nil {
			log.Fatal(err)
		}
	}
}
