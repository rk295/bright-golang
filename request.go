package bright

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/rk295/bright-golang/restclient"
)

var (
	restClient restclient.Client
)

func init() {
	restClient = &http.Client{}
}

func (c *Client) makeRequest(p string, target interface{}, params ...url.Values) error {

	// Make sure we have a valid token
	if err := c.performAuth(); err != nil {
		return err
	}

	u, err := url.Parse(apiURL)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, p)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	for _, value := range params {
		for k, v := range value {
			c.Logger.Debugf("k: %s v: %v", k, v)
			q.Add(k, v[0])
		}
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Set("applicationId", c.config.applicationID)

	if c.auth.token != "" {
		req.Header.Set("token", c.auth.token)
	}

	c.Logger.Debugf("making request to %s", req.URL.String())

	resp, err := restClient.Do(req)
	if err != nil {
		return err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.Logger.Trace("body: ", string(bodyBytes))

	c.Logger.Debugf("status code: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got status code %d", resp.StatusCode)
	}

	return json.Unmarshal(bodyBytes, &target)
}
