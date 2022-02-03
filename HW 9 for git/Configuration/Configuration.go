package Config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type ourSite struct {
	Url      string `json:"url"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Location string `json:"location"`
}

func Load_yaml() {

	text, err := os.ReadFile("Configuration/import.yaml")
	if err != nil {
		panic(err)
	}
	var site ourSite
	if err := yaml.Unmarshal(text, &site); err != nil {
		panic("Ошибка загрузки файла yaml.")
	}

	fmt.Println(&site)
	fmt.Println("Url = " + site.Url)
	fmt.Println("Host = " + site.Host)
	fmt.Println("User = " + site.User)
	fmt.Println("Location = " + site.Location)
}

func Load_json() {

	text, err := os.ReadFile("Configuration/import.json")
	if err != nil {
		panic(err)
	}
	var site ourSite
	if err := json.Unmarshal(text, &site); err != nil {
		panic("Ошибка загрузки файла json.")
	}

	fmt.Println(&site)
	fmt.Println("Url = " + site.Url)
	fmt.Println("Host = " + site.Host)
	fmt.Println("User = " + site.User)
	fmt.Println("Location = " + site.Location)

}
