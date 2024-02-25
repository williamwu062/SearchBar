package spotify

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	clientID     string `yaml:"client_id"`
	clientSecret string `yaml:"client_secret"`
}

func getSecrets(secrets_file string) (*Config, error) {
	file, err := os.Open(secrets_file)
	if err != nil {
		fmt.Println("Error opening file")
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Error decoding YAML:", err)
		return nil, err
	}

	return &config, nil
}

func connectAPI() {
	config, err := getSecrets("../secrets/yaml")
	if err != nil {
		fmt.Println("Error opening file")
	}

	url := "https://accounts.spotify.com/api/token"
	requestBody := []byte(fmt.Sprintf(`{"client_credential":%s, "client_secret":%s}`, config.clientID, config.clientSecret))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error processing request")
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Unexpected status code:", resp.StatusCode)
		return
	}

	var responseBody []byte
	_, err = resp.Body.Read(responseBody)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response:", string(responseBody))
}
