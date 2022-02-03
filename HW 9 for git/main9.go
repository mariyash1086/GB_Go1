package main

import (
	Config "Configuration/Configuration"
	"fmt"
)

func main() {

	fmt.Println("JSON:============================")
	Config.Load_json()
	fmt.Println("YAML:============================")
	Config.Load_yaml()

}
