package bright

import "fmt"

// GetElectricityConsumptionResourceID returns the resource ID of the
// electricity consumption resource.
func (c *Client) GetElectricityConsumptionResourceID() (string, error) {
	return c.findByClassifier(ElectricityConsumption)
}

// GetGasConsumptionResourceID returns the resource ID of the gas consumption
// resource.
func (c *Client) GetGasConsumptionResourceID() (string, error) {
	return c.findByClassifier(GasConsumption)
}

// GetElectricityConsumptionCostResourceID returns the resource ID of the
// electricity consumption cost resource.
func (c *Client) GetElectricityConsumptionCostResourceID() (string, error) {
	return c.findByClassifier(ElectricityConsumptionCost)
}

// GetGasConsumptionCostResourceID returns the resource ID of the gas
// consumption cost resource.
func (c *Client) GetGasConsumptionCostResourceID() (string, error) {
	return c.findByClassifier(GasConsumptionCost)
}

// GetElectricityCurrent returns a ResourceCurrent for the electricity
// consumption resource.
func (c *Client) GetElectricityCurrent() (ResourceCurrent, error) {
	return c.getCurrent(ElectricityConsumption)
}

// GetGasCurrent returns a ResourceCurrent for the gas consumption resource.
func (c *Client) GetGasCurrent() (ResourceCurrent, error) {
	return c.getCurrent(GasConsumption)
}

// GetElectricityCurrentWatts returns the current Watts for the electricity
// consumption resource.
func (c *Client) GetElectricityCurrentWatts() (int, error) {
	return c.getCurrentWatts(ElectricityConsumption)
}

// GetGasCurrentWatts returns the currents Watts for the gas consumption
// resource.
func (c *Client) GetGasCurrentWatts() (int, error) {
	return c.getCurrentWatts(GasConsumption)
}

func (c *Client) findByClassifier(cl ClassifierField) (string, error) {
	res, err := c.GetResources()
	if err != nil {
		return "", err
	}

	for _, r := range res {
		if r.Classifier == cl && r.DataSourceResourceTypeInfo.Unit == KWH {
			return r.ResourceID, nil
		}
	}
	return "", fmt.Errorf("failed to find any resources matching classifier id %s", cl)
}

func (c *Client) getCurrent(cl ClassifierField) (ResourceCurrent, error) {
	resourceID, err := c.findByClassifier(cl)
	if err != nil {
		return ResourceCurrent{}, err
	}
	return c.GetResourceCurrent(resourceID)
}

func (c *Client) getCurrentWatts(cl ClassifierField) (int, error) {
	current, err := c.getCurrent(cl)
	if err != nil {
		return 0, err
	}

	// Would love to know what the first field is!
	return current.Data[0][1], nil
}
