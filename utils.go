package bright

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

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

func readTestData(fileName string) (io.ReadCloser, error) {
	testFile, err := os.Open(path.Join(testDataDir, fileName))
	if err != nil {
		return nil, err
	}
	return ioutil.NopCloser(testFile), nil
}

func readBodyUnmarshall(body io.ReadCloser, target interface{}) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, &target)
}
