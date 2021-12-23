package bright

import "fmt"

func (c *Client) GetElectricityConsumptionResourceID() (string, error) {
	return c.findByResourceTypeID(electricityConsumptionResourceTypeId)
}

func (c *Client) GetGasConsumptionResourceID() (string, error) {
	return c.findByResourceTypeID(gasConsumptionResourceTypeId)
}

func (c *Client) GetElectricityConsumptionCostResourceID() (string, error) {
	return c.findByResourceTypeID(electricityConsumptionCostResourceTypeId)
}

func (c *Client) GetGasConsumptionCostResourceID() (string, error) {
	return c.findByResourceTypeID(gasConsumptionCostResourceTypeId)
}

func (c *Client) GetElectricityCurrent() (ResourceCurrent, error) {
	return c.getCurrent(electricityConsumptionResourceTypeId)
}

func (c *Client) GetGasCurrent() (ResourceCurrent, error) {
	return c.getCurrent(gasConsumptionResourceTypeId)
}

func (c *Client) GetElectricityCurrentKwh() (int, error) {
	return c.getCurrentKwh(electricityConsumptionResourceTypeId)
}

func (c *Client) GetGasCurrentKwh() (int, error) {
	return c.getCurrentKwh(gasConsumptionResourceTypeId)
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

func (c *Client) getCurrentKwh(typeID string) (int, error) {
	current, err := c.getCurrent(typeID)
	if err != nil {
		return 0, err
	}

	// Would love to know what the first field is!
	return current.Data[0][1], nil
}
