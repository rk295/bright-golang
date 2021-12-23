package bright

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

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

	c.Logger.Debugf("making request to %s", req.URL.String())

	return c.doRequest(req, target)
}

func setLogger(c *Client) {
	if c.Logger == nil {
		fmt.Println("making our own logger")
		c.Logger = logrus.New()
		if c.LogLevel != (logrus.Level(0)) {
			c.Logger.SetLevel(c.LogLevel)
		}
		c.Logger.SetFormatter(
			&logrus.TextFormatter{
				ForceColors:     true,
				FullTimestamp:   true,
				TimestampFormat: time.RFC3339Nano,
			},
		)
	}
}

func (c *Client) doRequest(r *http.Request, t interface{}) error {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("applicationId", c.config.applicationID)

	if c.auth.token != "" {
		r.Header.Set("token", c.auth.token)
	}

	client := &http.Client{}

	resp, err := client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	c.Logger.Debugf("status code: %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("got status code %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.Logger.Trace("body: ", string(bodyBytes))

	return json.Unmarshal(bodyBytes, &t)

}
