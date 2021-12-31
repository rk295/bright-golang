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

func TestFindByClassifier(t *testing.T) {

	tests := []struct {
		classifier ClassifierField
		fileName   string
		resourceID string
		unit       string
	}{
		{
			resourceID: "86923d05-3a4a-45dc-930c-a4fe355ef09b",
			fileName:   "resource-rk.json",
			classifier: ElectricityConsumption,
			unit:       KWH,
		},
		{
			resourceID: "7f8ab3b7-7285-4aaa-a0f4-1233cc52d75c",
			fileName:   "resource-rk.json",
			classifier: ElectricityConsumptionCost,
			unit:       Pence,
		},
		{
			resourceID: "5dc089fd-9e70-407b-8fc8-9c65f243d9f0",
			fileName:   "resource-rk.json",
			classifier: GasConsumption,
			unit:       KWH,
		},
		{
			resourceID: "f16983d0-8c56-4b3b-bb32-38066df6b744",
			fileName:   "resource-rk.json",
			classifier: GasConsumptionCost,
			unit:       Pence,
		},
	}

	c, err := NewTestClient()
	assert.Nil(t, err)

	for _, test := range tests {
		t.Logf("running test with %s", test.fileName)
		r, err := readTestData(test.fileName)
		assert.Nil(t, err)

		mock.GetDoFunc = mock.DefaultDo(r, http.StatusOK)
		resp, err := c.findByClassifier(test.classifier, test.unit)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, test.resourceID, resp)
	}
}

