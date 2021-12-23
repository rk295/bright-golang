package bright

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// NewClient returns a bright.Client.
func NewClient(c *Config) (*Client, error) {
	c.applicationID = applicationID
	client := &Client{config: c}
	return client, nil
}

// NewClientFromEnv returns a bright.Client drawing its configuration from the
// os environment.
func NewClientFromEnv() (*Client, error) {
	c := &Config{}

	username := os.Getenv(usernameEnv)
	if username == "" {
		return &Client{}, fmt.Errorf("%s is not set", usernameEnv)
	}
	c.Username = username

	password := os.Getenv(passwordEnv)
	if password == "" {
		return &Client{}, fmt.Errorf("%s is not set", passwordEnv)
	}
	c.Password = password

	client := &Client{config: c}
	return client, nil
}

// WithLogger configures a client to use an existing logsrus logger
func (c *Client) WithLogger(l *logrus.Logger) *Client {
	c.Logger = l
	setLogger(c)
	return c
}

// WithLevel configures the logger within a client to the given logging level
func (c *Client) WithLevel(level logrus.Level) *Client {
	c.Logger.SetLevel(level)
	return c
}
