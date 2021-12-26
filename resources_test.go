package bright

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rk295/bright-golang/mock"
)

func init() {
	restClient = &mock.MockClient{}
}

func TestGetResource(t *testing.T) {

	testFiles := []string{
		"resource-ce.json",
		"resource-rk.json",
	}

	c, err := NewTestClient()
	assert.Nil(t, err)

	for _, file := range testFiles {
		t.Logf("running test with %s", file)
		r, err := readTestData(file)
		assert.Nil(t, err)

		mock.GetDoFunc = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		}

		resp, err := c.GetResources()
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	}
}
