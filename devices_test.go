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

func TestGetDevices(t *testing.T) {

	testFiles := []string{
		"devices.json",
	}

	c, err := NewTestClient()
	assert.Nil(t, err)

	for _, file := range testFiles {
		t.Logf("running test with %s", file)
		r, err := readTestData(file)
		assert.Nil(t, err)

		mock.GetDoFunc = mock.DefaultDo(r, http.StatusOK)
		resp, err := c.GetDevices()
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	}
}
