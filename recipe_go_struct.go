package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"jsonutils"
)

const GoStructTemplate = `// File genereated by Configd.
package configd

import (
	"encoding/json"
	"strings"
)

// Contains your configuration
// data.
var Settings = ParseJSON(%s)

func ParseJSON(j string) (m Config) {
	sr := strings.NewReader(j)
	dec := json.NewDecoder(sr)

	err := dec.Decode(&m)
	
	if err != nil {
		panic(err)
	}

	return
}
`

func ExportForGoStructs(config string) {
	CreateVendorFolder()
	
	os.Remove("vendor/configd/structs.go")

	var m *jsonutils.Model

	data := GetConfig(config)

	m, err := jsonutils.FromBytes(data, "Config")
	if err != nil {
		log.Fatal(err)
	}

	// Write package file responsible for
	// decoding json data.
	config = fmt.Sprintf("`%s`", string(data))
	moduleStr := fmt.Sprintf(GoStructTemplate, config)
	module := []byte(moduleStr)
	err = ioutil.WriteFile("vendor/configd/c.go", module, 0700)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("vendor/configd/structs.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Fprintf(f,"package configd\n\n")

	m.Writer = f
	m.WriteGo()

}


