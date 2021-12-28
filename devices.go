package bright

// GetDevices returns all Devices known by the API.
func (c *Client) GetDevices() (Devices, error) {
	devices := &Devices{}
	err := c.makeRequest(deviceEndpoint, &devices)
	return *devices, err
}
