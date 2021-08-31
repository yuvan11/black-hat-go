package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telnet"`
	Plan         string `json:"plan"`
	Https        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}

func (client *Client) APIInfo() (*APIInfo, error) {

	resp, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", apiURL, client.ApiKey))

	if err != nil {
		fmt.Println("Error in connecting client URL")
		return nil, nil
	}

	defer resp.Body.Close()

	var ret APIInfo

	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
