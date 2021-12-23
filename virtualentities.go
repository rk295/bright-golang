package bright

import "path"

// GetVEs returns all the Virtual Entities known to the API.
func (c *Client) GetVEs() (VirtualEntities, error) {
	virtualEntities := &VirtualEntities{}
	err := c.makeRequest(veEndpoint, virtualEntities)
	return *virtualEntities, err
}

// GetVE returns a single Virtual Entity as specified by the virtual entity id
// ve.
func (c *Client) GetVE(ve string) (VirtualEntity, error) {
	virtualEntity := &VirtualEntity{}
	err := c.makeRequest(path.Join(veEndpoint, ve), virtualEntity)
	return *virtualEntity, err
}
