package bright

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"time"
)

func (c *Client) performAuth() error {

	// Auth if token is empty
	if c.auth.token == "" {
		c.Logger.Debug("token is empty, authenticating")
		return c.getToken()
	}

	// Auth if token expiry is in the past
	t := time.Now()
	if c.auth.expiry.Before(t) {
		c.Logger.Debugf("expiry (%s) is in the past (now: %s)\n", c.auth.expiry, t)
		return c.getToken()
	}
	return nil
}

func (c *Client) getToken() error {

	u, err := url.Parse(apiURL)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, authEndpoint)

	c.Logger.Debugf("making auth request to url: %s", u)

	authRequest := &AuthRequest{
		Username:      c.config.Username,
		Password:      c.config.Password,
		ApplicationID: c.config.applicationID,
	}

	jsonPayload, err := json.Marshal(authRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	r := &AuthResponse{}
	if err := c.doRequest(req, r); err != nil {
		return err
	}

	c.auth.token = r.Token
	c.auth.expiry = time.Unix(r.Exp, 0)
	c.Logger.Debugf("got token, expires at %s", c.auth.expiry)
	return nil
}
