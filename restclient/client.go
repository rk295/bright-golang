package restclient

import "net/http"

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

func Do(r *http.Request) (*http.Response, error) {
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(r)
	if err != nil {
		return &http.Response{}, err
	}
	return resp, nil
}
