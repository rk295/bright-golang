package bright

// GetResourceTypes returns all ResourceTypes known by the API.
func (c *Client) GetResourceTypes() (ResourceTypes, error) {
	reTypeList := &ResourceTypes{}
	err := c.makeRequest(resourceTypeEndpoint, &reTypeList)
	return *reTypeList, err
}
