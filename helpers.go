package bright

import "fmt"

// GetElectricityConsumptionResourceID returns the resource ID of the
// electricity consumption resource.
func (c *Client) GetElectricityConsumptionResourceID() (string, error) {
	return c.findByResourceTypeID(electricityConsumptionResourceTypeId)
}

// GetGasConsumptionResourceID returns the resource ID of the gas consumption
// resource.
func (c *Client) GetGasConsumptionResourceID() (string, error) {
	return c.findByResourceTypeID(gasConsumptionResourceTypeId)
}

// GetElectricityConsumptionCostResourceID returns the resource ID of the
// electricity consumption cost resource.
func (c *Client) GetElectricityConsumptionCostResourceID() (string, error) {
	return c.findByResourceTypeID(electricityConsumptionCostResourceTypeId)
}

// GetGasConsumptionCostResourceID returns the resource ID of the gas
// consumption cost resource.
func (c *Client) GetGasConsumptionCostResourceID() (string, error) {
	return c.findByResourceTypeID(gasConsumptionCostResourceTypeId)
}

// GetElectricityCurrent returns a ResourceCurrent for the electricity
// consumption resource.
func (c *Client) GetElectricityCurrent() (ResourceCurrent, error) {
	return c.getCurrent(electricityConsumptionResourceTypeId)
}

// GetGasCurrent returns a ResourceCurrent for the gas consumption resource.
func (c *Client) GetGasCurrent() (ResourceCurrent, error) {
	return c.getCurrent(gasConsumptionResourceTypeId)
}

// GetElectricityCurrentWatts returns the current Watts for the electricity
// consumption resource.
func (c *Client) GetElectricityCurrentWatts() (int, error) {
	return c.getCurrentWatts(electricityConsumptionResourceTypeId)
}

// GetGasCurrentWatts returns the currents Watts for the gas consumption
// resource.
func (c *Client) GetGasCurrentWatts() (int, error) {
	return c.getCurrentWatts(gasConsumptionResourceTypeId)
}

func (c *Client) findByResourceTypeID(typeId string) (string, error) {
	res, err := c.GetResources()
	if err != nil {
		return "", err
	}

	for _, r := range res {
		if r.ResourceTypeID == typeId {
			return r.ResourceID, nil
		}
	}
	return "", fmt.Errorf("failed to find any resources matching type id %s", typeId)
}

func (c *Client) getCurrent(typeID string) (ResourceCurrent, error) {
	resourceID, err := c.findByResourceTypeID(typeID)
	if err != nil {
		return ResourceCurrent{}, err
	}
	return c.GetResourceCurrent(resourceID)
}

func (c *Client) getCurrentWatts(typeID string) (int, error) {
	current, err := c.getCurrent(typeID)
	if err != nil {
		return 0, err
	}

	// Would love to know what the first field is!
	return current.Data[0][1], nil
}
