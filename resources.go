package bright

import (
	"net/url"
	"path"
	"time"
)

// GetResources returns all Resources known by the API.
func (c *Client) GetResources() (Resources, error) {
	reList := &Resources{}
	err := c.makeRequest(resourceEndpoint, reList)
	return *reList, err
}

// GetResource returns a single resource identified by the resource ID re.
func (c *Client) GetResource(re string) (Resource, error) {
	resource := &Resource{}
	err := c.makeRequest(path.Join(resourceEndpoint, re), resource)
	return *resource, err
}

// GetResourceCurrent returns the current usage of the Resource re
func (c *Client) GetResourceCurrent(re string) (ResourceCurrent, error) {
	current := &ResourceCurrent{}
	err := c.makeRequest(path.Join(resourceEndpoint, re, currentEndpoint), current)
	return *current, err
}

// GetResourceReading returns a Reading for the resource re in the range from-to
// applying the aggregation function  and the period specified.
func (c *Client) GetResourceReading(re, period, function string, from, to time.Time) (Reading, error) {
	r := &Reading{}

	params := url.Values{}
	params.Add("period", period)
	params.Add("function", function)
	params.Add("from", from.Format(readingDateFormatStr))
	params.Add("to", to.Format(readingDateFormatStr))

	err := c.makeRequest(path.Join(resourceEndpoint, re, readingEndpoint), r, params)
	return *r, err
}
