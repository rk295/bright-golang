package bright

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Config struct {
	Username      string
	Password      string
	applicationID string
}

type Client struct {
	config   *Config
	auth     Auth
	Logger   *logrus.Logger
	LogLevel logrus.Level
}

type Auth struct {
	token  string
	expiry time.Time
}

type ClassifierField string

const (
	ElectricityConsumption     ClassifierField = "electricity.consumption"
	ElectricityConsumptionCost ClassifierField = "electricity.consumption.cost"
	GasConsumption             ClassifierField = "gas.consumption"
	GasConsumptionCost         ClassifierField = "gas.consumption.cost"
)

type UnitField string

const (
	KwH   UnitField = "kWh"
	Pence UnitField = "pence"
)

type TypeField string

const (
	Gas         TypeField = "GAS"
	Electricity TypeField = "ELEC"
	Power       TypeField = "PWER"
)

type AuthRequest struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	ApplicationID string `json:"applicationId"`
}

type AuthResponse struct {
	Valid                   bool          `json:"valid"`
	Token                   string        `json:"token"`
	Exp                     int64         `json:"exp"`
	UserGroups              []interface{} `json:"userGroups"`
	FunctionalGroupAccounts []interface{} `json:"functionalGroupAccounts"`
	AccountID               string        `json:"accountId"`
	IsTempAuth              bool          `json:"isTempAuth"`
	Name                    string        `json:"name"`
}

type VirtualEntities []VirtualEntity

type VirtualEntity struct {
	Clone         bool          `json:"clone"`
	Active        bool          `json:"active"`
	ApplicationID string        `json:"applicationId"`
	PostalCode    string        `json:"postalCode"`
	Resources     []VEResources `json:"resources"`
	OwnerID       string        `json:"ownerId"`
	Name          string        `json:"name"`
	VeChildren    []interface{} `json:"veChildren"`
	VeTypeID      string        `json:"veTypeId"`
	VeID          string        `json:"veId"`
	UpdatedAt     time.Time     `json:"updatedAt"`
	CreatedAt     time.Time     `json:"createdAt"`
}

type VEResources struct {
	ResourceID     string `json:"resourceId"`
	ResourceTypeID string `json:"resourceTypeId"`
	Name           string `json:"name"`
}

type Resources []Resource

type Resource struct {
	Active                     bool                       `json:"active"`
	ResourceTypeID             string                     `json:"resourceTypeId"`
	OwnerID                    string                     `json:"ownerId"`
	Name                       string                     `json:"name"`
	Description                string                     `json:"description"`
	Label                      string                     `json:"label"`
	DataSourceResourceTypeInfo DataSourceResourceTypeInfo `json:"dataSourceResourceTypeInfo"`
	DataSourceType             string                     `json:"dataSourceType"`
	Classifier                 ClassifierField            `json:"classifier"`
	BaseUnit                   UnitField                  `json:"baseUnit"`
	ResourceID                 string                     `json:"resourceId"`
	UpdatedAt                  time.Time                  `json:"updatedAt"`
	CreatedAt                  time.Time                  `json:"createdAt"`
	DataSourceUnitInfo         DataSourceUnitInfo         `json:"dataSourceUnitInfo"`
}
type DataSourceResourceTypeInfo struct {
	Unit UnitField `json:"unit"`
	Type TypeField `json:"type"`
}
type DataSourceUnitInfo struct {
	Shid string `json:"shid"`
}

type ResourceCurrent struct {
	Status         string          `json:"status"`
	Name           string          `json:"name"`
	ResourceTypeID string          `json:"resourceTypeId"`
	ResourceID     string          `json:"resourceId"`
	Data           [][]int         `json:"data"`
	Units          string          `json:"units"`
	Classifier     ClassifierField `json:"classifier"`
}

type Readings struct {
	Status         string          `json:"status"`
	Name           string          `json:"name"`
	ResourceTypeID string          `json:"resourceTypeId"`
	ResourceID     string          `json:"resourceId"`
	Query          Query           `json:"query"`
	Data           [][]float32     `json:"data"`
	Units          UnitField       `json:"units"`
	Classifier     ClassifierField `json:"classifier"`
}

type Query struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Period   string `json:"period"`
	Function string `json:"function"`
}
