package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const PHPTemplate = `<?php

$configdString = '%s';

$configd = json_decode($configdString);

?>`

func ExportForPHP(config string) {

	data := GetConfig(config)
	config = string(data)
	moduleStr := fmt.Sprintf(PHPTemplate, config)
	module := []byte(moduleStr)
	err := ioutil.WriteFile("configd.php", module, 0700)
	if err != nil {
		log.Fatal(err)
	}
}


