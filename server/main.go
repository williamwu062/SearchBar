package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	clientID     string `yaml:"client_id"`
	clientSecret string `yaml:client_secret`
}

func main() {
	file, err := os.Open("../secrets.yml")
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Error decoding YAML:", err)
		return
	}
}
