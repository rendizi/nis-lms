package reqs

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Make requests to NIS Mektep
type AuthResponse struct {
	AccessToken string `json:"accessToken"`
	// Add other fields from the response if needed
}

func Login(login, password string) (string, error) {
	payload := map[string]interface{}{
		"action":      "v1/Users/Authenticate",
		"operationId": "03b5c3d1-bb48-4e97-a378-62ecf66d6c11",
		"username":    login,
		"password":    password,
		"deviceInfo":  "SM-G977N",
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Prepare the request
	req, err := http.NewRequest("POST", "https://identity.micros.nis.edu.kz/v1/Users/Authenticate", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "Culture=ru-RU; lang=ru-RU")
	req.Header.Set("Host", "identity.micros.nis.edu.kz")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return "", err
	}

	var authResponse AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return "", err
	}

	return authResponse.AccessToken, nil
}
