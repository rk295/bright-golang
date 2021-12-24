package bright

import (
	"time"

	"github.com/sirupsen/logrus"
)

// Auth holds the authentication token and expiry time.
type Auth struct {
	token  string
	expiry time.Time
}

// AuthRequest represents a request to the Bright API to authenticate ourselves.
type AuthRequest struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	ApplicationID string `json:"applicationId"`
}

// AuthResponse holds the response from an authentication response.
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

// Config represents the configuration of the client. Both Username and Password
// must be set, or if using NewClientFromEnv they will be set for you using the
// os environment.
type Config struct {
	Username      string
	Password      string
	applicationID string
}

// Client represents a bright API client. You can optionally add an existing
// logger with WithLogger() and set its level with WithLevel().
type Client struct {
	config   *Config
	auth     Auth
	Logger   *logrus.Logger
	LogLevel logrus.Level
}

// ClassfierField is the type applied for the classifier fields returned by the
// API.
type ClassifierField string

const (
	ElectricityConsumption     ClassifierField = "electricity.consumption"
	ElectricityConsumptionCost ClassifierField = "electricity.consumption.cost"
	GasConsumption             ClassifierField = "gas.consumption"
	GasConsumptionCost         ClassifierField = "gas.consumption.cost"
)

// DataSourceUnitField is the type used for the Units returned as part of a Resource.
type DataSourceUnitField string

const (
	KwH   DataSourceUnitField = "kWh"
	Pence DataSourceUnitField = "pence"
)

// DataSourceResourceTypeInfo holds the Unit and Type of the Data Source
// Resource returned by the API.
type DataSourceResourceTypeInfo struct {
	Unit DataSourceUnitField `json:"unit"`
	Type TypeField           `json:"type"`
}

// DataSourceUnitInfo holds Data Source unit info.
type DataSourceUnitInfo struct {
	Shid string `json:"shid"`
}

// Query represents a query to the readings API.
type Query struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Period   string `json:"period"`
	Function string `json:"function"`
}

// Reading represents a single reading returned from the API.
type Reading struct {
	Status         string           `json:"status"`
	Name           string           `json:"name"`
	ResourceTypeID string           `json:"resourceTypeId"`
	ResourceID     string           `json:"resourceId"`
	Query          Query            `json:"query"`
	Data           [][]float32      `json:"data"`
	Units          ReadingUnitField `json:"units"`
	Classifier     ClassifierField  `json:"classifier"`
}

// Resources is a slice of Resources.
type Resources []Resource

// Resource represents a resource as defined by the API.
type Resource struct {
	Active                     bool                       `json:"active"`
	ResourceTypeID             TypeIDField                `json:"resourceTypeId"`
	OwnerID                    string                     `json:"ownerId"`
	Name                       string                     `json:"name"`
	Description                string                     `json:"description"`
	Label                      string                     `json:"label"`
	DataSourceResourceTypeInfo DataSourceResourceTypeInfo `json:"dataSourceResourceTypeInfo"`
	DataSourceType             string                     `json:"dataSourceType"`
	Classifier                 ClassifierField            `json:"classifier"`
	BaseUnit                   DataSourceUnitField        `json:"baseUnit"`
	ResourceID                 string                     `json:"resourceId"`
	UpdatedAt                  time.Time                  `json:"updatedAt"`
	CreatedAt                  time.Time                  `json:"createdAt"`
	DataSourceUnitInfo         DataSourceUnitInfo         `json:"dataSourceUnitInfo"`
}

// ReadingUnitField represents all the possible units the API can return.
type ReadingUnitField string

const (
	Watts       ReadingUnitField = "W"
	MetersCubed ReadingUnitField = "m3"
)

// Resourcecurrent represents the current (as in now) usage of the Resource.
type ResourceCurrent struct {
	Status         string           `json:"status"`
	Name           string           `json:"name"`
	ResourceTypeID TypeIDField      `json:"resourceTypeId"`
	ResourceID     string           `json:"resourceId"`
	Data           [][]int          `json:"data"`
	Units          ReadingUnitField `json:"units"`
	Classifier     ClassifierField  `json:"classifier"`
}

// TypeIDField holds the various type IDs of resources.
type TypeIDField string

const (
	ElectricityConsumptionResource     TypeIDField = "e3a5db34-6e0c-4221-9653-8d33e27511ba"
	ElectricityConsumptionCostResource TypeIDField = "78859e39-611e-4e84-a402-1d4460abcb56"
	GasConsumptionResource             TypeIDField = "08ab415f-d851-423f-adf4-c2b1e0529e27"
	GasConsumptionCostResource         TypeIDField = "a6b95f41-771d-4bd2-99f4-93ee43c38f5a"
)

// TypeField is the type used for the type in the DataSourceResourceTypeInfo
// struct.
type TypeField string

const (
	Gas         TypeField = "GAS"
	Electricity TypeField = "ELEC"
	Power       TypeField = "PWER"
)

// VirtualEntities is a slice of VirtualEntities
type VirtualEntities []VirtualEntity

// VirtualEntity represents a single Virtual Entity returned by the API.
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

// VEResources represents the Resources returned as part of a VirtualEntity.
type VEResources struct {
	ResourceID     string `json:"resourceId"`
	ResourceTypeID string `json:"resourceTypeId"`
	Name           string `json:"name"`
}
