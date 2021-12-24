package bright

import "fmt"

// GetElectricityConsumptionResourceID returns the resource ID of the
// electricity consumption resource.
func (c *Client) GetElectricityConsumptionResourceID() (string, error) {
	return c.findByResourceTypeID(ElectricityConsumptionResource)
}

// GetGasConsumptionResourceID returns the resource ID of the gas consumption
// resource.
func (c *Client) GetGasConsumptionResourceID() (string, error) {
	return c.findByResourceTypeID(GasConsumptionResource)
}

// GetElectricityConsumptionCostResourceID returns the resource ID of the
// electricity consumption cost resource.
func (c *Client) GetElectricityConsumptionCostResourceID() (string, error) {
	return c.findByResourceTypeID(ElectricityConsumptionCostResource)
}

// GetGasConsumptionCostResourceID returns the resource ID of the gas
// consumption cost resource.
func (c *Client) GetGasConsumptionCostResourceID() (string, error) {
	return c.findByResourceTypeID(GasConsumptionCostResource)
}

// GetElectricityCurrent returns a ResourceCurrent for the electricity
// consumption resource.
func (c *Client) GetElectricityCurrent() (ResourceCurrent, error) {
	return c.getCurrent(ElectricityConsumptionResource)
}

// GetGasCurrent returns a ResourceCurrent for the gas consumption resource.
func (c *Client) GetGasCurrent() (ResourceCurrent, error) {
	return c.getCurrent(GasConsumptionResource)
}

// GetElectricityCurrentWatts returns the current Watts for the electricity
// consumption resource.
func (c *Client) GetElectricityCurrentWatts() (int, error) {
	return c.getCurrentWatts(ElectricityConsumptionResource)
}

// GetGasCurrentWatts returns the currents Watts for the gas consumption
// resource.
func (c *Client) GetGasCurrentWatts() (int, error) {
	return c.getCurrentWatts(GasConsumptionResource)
}

func (c *Client) findByResourceTypeID(typeId TypeIDField) (string, error) {
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

func (c *Client) getCurrent(typeID TypeIDField) (ResourceCurrent, error) {
	resourceID, err := c.findByResourceTypeID(typeID)
	if err != nil {
		return ResourceCurrent{}, err
	}
	return c.GetResourceCurrent(resourceID)
}

func (c *Client) getCurrentWatts(typeID TypeIDField) (int, error) {
	current, err := c.getCurrent(typeID)
	if err != nil {
		return 0, err
	}

	// Would love to know what the first field is!
	return current.Data[0][1], nil
}
