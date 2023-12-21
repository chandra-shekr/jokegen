package jokegen

import (
	"encoding/json"
	"io"
	"net/http"
)

type response struct {
	Delivery string `json:"delivery"`
}

func GenRandomJoke() (string, error) {

	resp, err := http.Get("https://v2.jokeapi.dev/joke/Any")

	if err != nil {

		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result response

	if err := json.Unmarshal(body, &result); err != nil {

		return "", err
	}

	return result.Delivery, nil

}
