package main

import (
	"fmt"
	"github.com/mariyash1086/GB_Level1/Configuration"
)

func main() {

	config, err := Configuration.Load()
	fmt.Println(config)
	fmt.Println(err)
}
