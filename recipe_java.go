package main

import (
	"encoding/json"
	"jsonutils"
	"log"
	"os"
)

func ExportForJava(config string) {

	var m *jsonutils.Model

	data := GetConfig(config)

	m, err := jsonutils.FromBytes(data, "Config")
	if err != nil {
		log.Fatal(err)
	}

	CreatePackage()

	f, err := os.Create("app/settings/Config.java")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m.Writer = f
	m.WriteJava()
}

func GetConfig(config string) []byte {
	b := []byte(config)
	i, err := jsonutils.ParseJson(b)
	if err != nil {
		log.Fatal(err)
	}

	configMap, ok := i.(map[string]interface{})

	if !ok {
		log.Fatal("Could not read configuration data.")
	}

	b, err = json.Marshal(configMap["Config"])
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func CreatePackage() {
	if _, err := os.Stat("app/settings"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err = os.MkdirAll("app/settings", 0700)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		os.Remove("app/settings/Config.java")
	}

}
