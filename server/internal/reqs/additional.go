package reqs

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type additional struct {
	Klass  string `json:"klass"`
	School struct {
		Name struct {
			En string `json:"en"`
		} `json:"name"`
	} `json:"school"`
}

func AdditionalInfo(token string) (string, string, error) {
	payload := map[string]interface{}{
		"applicationType": "ContingentAPI",
		"action":          "Api/AdditionalUserInfo",
		"operationId":     "7ea4e116-5585-4518-b180-453815338986",
		"token":           token,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", "", err
	}

	// Prepare the request
	req, err := http.NewRequest("POST", "https://contingent.micros.nis.edu.kz/Api/AdditionalUserInfo", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "Culture=ru-RU; lang=ru-RU")
	req.Header.Set("Host", "contingent.micros.nis.edu.kz")
	req.Header.Set("authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()
	if resp.Status != "200 OK" {
		return "", "", err
	}

	var additionalInfo additional
	err = json.NewDecoder(resp.Body).Decode(&additionalInfo)
	if err != nil {
		return "", "", err
	}

	return additionalInfo.Klass, additionalInfo.School.Name.En, nil
}
