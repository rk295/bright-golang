package bright

import (
	"net/url"
	"path"
	"time"
)

func (c *Client) GetResources() (Resources, error) {
	reList := &Resources{}
	err := c.makeRequest(resourceEndpoint, reList)
	return *reList, err
}

func (c *Client) GetResource(re string) (Resource, error) {
	resource := &Resource{}
	err := c.makeRequest(path.Join(resourceEndpoint, re), resource)
	return *resource, err
}

func (c *Client) GetResourceCurrent(re string) (ResourceCurrent, error) {
	current := &ResourceCurrent{}
	err := c.makeRequest(path.Join(resourceEndpoint, re, currentEndpoint), current)
	return *current, err
}

func (c *Client) GetResourceReadings(re, period, function string, from, to time.Time) (Readings, error) {
	r := &Readings{}

	params := url.Values{}
	params.Add("period", period)
	params.Add("function", function)
	params.Add("from", from.Format(readingsDateFormatStr))
	params.Add("to", to.Format(readingsDateFormatStr))

	err := c.makeRequest(path.Join(resourceEndpoint, re, readingsEndpoint), r, params)
	return *r, err
}
