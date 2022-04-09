package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

var apiUrl = "https://api.dev.runpod.io/graphql"

type Input struct {
	Query     string            `json:"query"`
	Variables map[string]string `json:"variables"`
}

func Query(input Input) (res *http.Response, err error) {
	jsonValue, err := json.Marshal(input)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", apiUrl+"?api_key="+viper.GetString("apiKey"), bytes.NewBuffer(jsonValue))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Second * 5}
	return client.Do(req)
}
