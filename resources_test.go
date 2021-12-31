package bright

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/rk295/bright-golang/mock"
)

func init() {
	restClient = &mock.MockClient{}
}

func TestGetResource(t *testing.T) {
	resourceTestFiles := []string{
		"resource-rk.json",
		"resource-ce.json",
	}

	c, err := NewTestClient()
	assert.Nil(t, err)

	for _, file := range resourceTestFiles {
		t.Logf("running test with %s", file)
		r, err := readTestData(file)
		assert.Nil(t, err)

		mock.GetDoFunc = mock.DefaultDo(r, http.StatusOK)
		resp, err := c.GetResources()
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	}
}

func TestGetResourceCurrent(t *testing.T) {

	testFiles := map[string]string{
		"86923d05-3a4a-45dc-930c-a4fe355ef09b": "resource-rk-86923d05-3a4a-45dc-930c-a4fe355ef09b.json",
	}

	c, err := NewTestClient()
	assert.Nil(t, err)

	for id, file := range testFiles {
		t.Logf("running test with %s", file)
		r, err := readTestData(file)
		assert.Nil(t, err)

		mock.GetDoFunc = mock.DefaultDo(r, http.StatusOK)
		resp, err := c.GetResourceCurrent(id)
		t.Log(err)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	}
}

func TestGetResourceReading(t *testing.T) {

	tests := []struct {
		resourceID string
		fileName   string
		period     string
		function   string
		to         time.Time
		from       time.Time
	}{
		{
			resourceID: "86923d05-3a4a-45dc-930c-a4fe355ef09b",
			fileName:   "resource-reading-rk-86923d05-3a4a-45dc-930c-a4fe355ef09b.json",
			period:     "PT30M",
			function:   "sum",
			from:       time.Date(2020, 11, 01, 00, 00, 00, 0, time.UTC),
			to:         time.Date(2020, 11, 10, 23, 59, 59, 0, time.UTC),
		},
	}

	c, err := NewTestClient()
	assert.Nil(t, err)

	for _, test := range tests {
		t.Logf("running test with %s", test.fileName)
		r, err := readTestData(test.fileName)
		assert.Nil(t, err)

		mock.GetDoFunc = mock.DefaultDo(r, http.StatusOK)
		resp, err := c.GetResourceReading(test.resourceID,
			test.period,
			test.function,
			test.from,
			test.to,
		)

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "2020-11-01T00:00:00", resp.Query.From)
		assert.Equal(t, "2020-11-10T23:59:59", resp.Query.To)
		assert.Equal(t, "PT30M", resp.Query.Period)
		assert.Equal(t, "sum", resp.Query.Function)
		assert.Equal(t, 480, len(resp.Data))
	}
}
