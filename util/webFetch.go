package util

import (
	"fmt"
	"io"
	"net/http"
)

// FetchWebPage retrieves the content of a web page specified by the URL.
// It includes a session cookie for Advent of Code authentication.
func FetchWebPage(url, sessionCookie string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Add the session cookie to the request
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionCookie))

	// Make the request
	response, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
