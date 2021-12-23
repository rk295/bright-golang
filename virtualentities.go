package bright

import "path"

func (c *Client) GetVEs() (VirtualEntities, error) {
	virtualEntities := &VirtualEntities{}
	err := c.makeRequest(veEndpoint, virtualEntities)
	return *virtualEntities, err
}

func (c *Client) GetVE(ve string) (VirtualEntity, error) {
	virtualEntity := &VirtualEntity{}
	err := c.makeRequest(path.Join(veEndpoint, ve), virtualEntity)
	return *virtualEntity, err
}
