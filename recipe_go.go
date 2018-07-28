package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const GoTemplate = `// File genereated by Configd.
package configd

import (
	"encoding/json"
	"strings"
	"log"
)

type dictionary map[string]interface{}

var Settings = ParseJSON(%s)

func Int(i interface{}) int64 {
	return i.(int64)
}

func Bool(i interface{}) bool {
	return i.(bool)
}

func String(i interface{}) string {
	return i.(string)
}

func Float64(i interface{}) float64 {
	return i.(float64)
}

func Map(i interface{}) map[string]interface{} {
	return i.(dictionary)
}

func ParseJSON(j string) (dictionary) {
	var d dictionary

	decoder := json.NewDecoder(strings.NewReader(j))

	err := decoder.Decode(&d)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return d
}

`

func ExportForGo(config string) {
	CreateVendorFolder()
	data := GetConfig(config)
	config = fmt.Sprintf("`%s`", string(data))
	moduleStr := fmt.Sprintf(GoTemplate, config)
	module := []byte(moduleStr)
	err := ioutil.WriteFile("vendor/configd/c.go", module, 0700)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateVendorFolder() {
	if _, err := os.Stat("vendor/configd"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err = os.MkdirAll("vendor/configd", 0700)
		if err != nil {
			log.Fatal(err)
		}
	}

}
