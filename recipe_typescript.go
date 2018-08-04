package main


import (
	"jsonutils"
	"log"
	"os"
)

func ExportForTS(config string) {

	var m *jsonutils.Model

	data := GetConfig(config)

	m, err := jsonutils.FromBytes(data, "Config")
	if err != nil {
		log.Fatal(err)
	}

	CreateModuleFolder()

	f, err := os.Create("configd/c.ts")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m.Writer = f
	m.WriteTypeScript()
}



func CreateModuleFolder() {
	if _, err := os.Stat("configd"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err = os.MkdirAll("configd", 0700)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		os.Remove("configd/c.ts")
	}

}
