package generator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type response struct {
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
}

func GenRandomJoke() (string, error) {

	resp, err := http.Get("https://v2.jokeapi.dev/joke/Any")

	if err != nil {

		return "", err
	}
	defer resp.Body.Close()

	if body, err := io.ReadAll(resp.Body); err != nil {

		return "", err
	} else {
		var result response
		if err := json.Unmarshal(body, &result); err != nil {

			return "", err
		}

		return fmt.Sprintf("%v\n%v", result.Setup, result.Delivery), nil

	}

}
