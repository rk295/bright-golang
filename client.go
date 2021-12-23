package bright

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func NewClient(c *Config) (*Client, error) {
	c.applicationID = applicationID
	client := &Client{config: c}
	return client, nil
}

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

func (c *Client) WithLogger(l *logrus.Logger) *Client {
	c.Logger = l
	setLogger(c)
	return c
}

func (c *Client) WithLevel(level logrus.Level) *Client {
	c.Logger.SetLevel(level)
	return c
}
