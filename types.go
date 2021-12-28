package bright

import (
	"time"

	"github.com/sirupsen/logrus"
)

// Auth holds the authentication token and expiry time.
type Auth struct {
	expiry time.Time
	token  string
}

// AuthRequest represents a request to the Bright API to authenticate ourselves.
type AuthRequest struct {
	ApplicationID string `json:"applicationId"`
	Password      string `json:"password"`
	Username      string `json:"username"`
}

// AuthResponse holds the response from an authentication response.
type AuthResponse struct {
	AccountID               string        `json:"accountId"`
	Exp                     int64         `json:"exp"`
	FunctionalGroupAccounts []interface{} `json:"functionalGroupAccounts"`
	IsTempAuth              bool          `json:"isTempAuth"`
	Name                    string        `json:"name"`
	Token                   string        `json:"token"`
	UserGroups              []interface{} `json:"userGroups"`
	Valid                   bool          `json:"valid"`
}

// Config represents the configuration of the client. Both Username and Password
// must be set, or if using NewClientFromEnv they will be set for you using the
// os environment.
type Config struct {
	applicationID string
	Password      string
	Username      string
}

// Client represents a bright API client. You can optionally add an existing
// logger with WithLogger() and set its level with WithLevel().
type Client struct {
	auth     Auth
	config   *Config
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

// DataSourceResourceTypeInfo holds the Unit and Type of the Data Source
// Resource returned by the API.
type DataSourceResourceTypeInfo struct {
	IsCost bool   `json:"isCost,omitempty"`
	Method string `json:"method,omitempty"`
	Range  string `json:"range,omitempty"`
	Type   string `json:"type,omitempty"`
	Unit   string `json:"unit,omitempty"`
}

// DataSourceUnitInfo holds Data Source unit info.
type DataSourceUnitInfo struct {
	Shid string `json:"shid"`
}

// Query represents a query to the readings API.
type Query struct {
	From     string `json:"from"`
	Function string `json:"function"`
	Period   string `json:"period"`
	To       string `json:"to"`
}

// Reading represents a single reading returned from the API.
type Reading struct {
	Classifier     ClassifierField  `json:"classifier"`
	Data           [][]float32      `json:"data"`
	Name           string           `json:"name"`
	Query          Query            `json:"query"`
	ResourceID     string           `json:"resourceId"`
	ResourceTypeID string           `json:"resourceTypeId"`
	Status         string           `json:"status"`
	Units          ReadingUnitField `json:"units"`
}

// Resources is a slice of Resources.
type Resources []Resource

// Resource represents a resource as defined by the API.
type Resource struct {
	Active                     bool                       `json:"active"`
	BaseUnit                   string                     `json:"baseUnit"`
	Classifier                 ClassifierField            `json:"classifier"`
	CreatedAt                  time.Time                  `json:"createdAt"`
	DataSourceResourceTypeInfo DataSourceResourceTypeInfo `json:"dataSourceResourceTypeInfo,omitempty"`
	DataSourceType             string                     `json:"dataSourceType"`
	DataSourceUnitInfo         DataSourceUnitInfo         `json:"dataSourceUnitInfo"`
	Description                string                     `json:"description"`
	Label                      string                     `json:"label"`
	Name                       string                     `json:"name"`
	OwnerID                    string                     `json:"ownerId"`
	ResourceID                 string                     `json:"resourceId"`
	ResourceTypeID             string                     `json:"resourceTypeId"`
	UpdatedAt                  time.Time                  `json:"updatedAt"`
}

// ReadingUnitField represents all the possible units the API can return.
type ReadingUnitField string

const (
	MetersCubed ReadingUnitField = "m3"
	Watts       ReadingUnitField = "W"
)

// Resourcecurrent represents the current (as in now) usage of the Resource.
type ResourceCurrent struct {
	Classifier     ClassifierField  `json:"classifier"`
	Data           [][]int          `json:"data"`
	Name           string           `json:"name"`
	ResourceID     string           `json:"resourceId"`
	ResourceTypeID string           `json:"resourceTypeId"`
	Status         string           `json:"status"`
	Units          ReadingUnitField `json:"units"`
}

// VirtualEntities is a slice of VirtualEntities
type VirtualEntities []VirtualEntity

// VirtualEntity represents a single Virtual Entity returned by the API.
type VirtualEntity struct {
	Active        bool          `json:"active"`
	ApplicationID string        `json:"applicationId"`
	Clone         bool          `json:"clone"`
	CreatedAt     time.Time     `json:"createdAt"`
	Name          string        `json:"name"`
	OwnerID       string        `json:"ownerId"`
	PostalCode    string        `json:"postalCode"`
	Resources     []VEResources `json:"resources"`
	UpdatedAt     time.Time     `json:"updatedAt"`
	VeChildren    []interface{} `json:"veChildren"`
	VeID          string        `json:"veId"`
	VeTypeID      string        `json:"veTypeId"`
}

// VEResources represents the Resources returned as part of a VirtualEntity.
type VEResources struct {
	Name           string `json:"name"`
	ResourceID     string `json:"resourceId"`
	ResourceTypeID string `json:"resourceTypeId"`
}

// ResourceTypes
type ResourceTypes []ResourceType

// ResourceType represents a single resource type.
// DataSourceResourceTypeInfo has to be an interface because one of the
// returned fields is a string not a object that can be Unmarshalled into
// DataSourceResourceTypeInfo
type ResourceType struct {
	Active                     bool              `json:"active"`
	BaseUnit                   string            `json:"baseUnit,omitempty"`
	Classifier                 ClassifierField   `json:"classifier,omitempty"`
	DataSourceResourceTypeInfo interface{}       `json:"dataSourceResourceTypeInfo"`
	DataSourceType             string            `json:"dataSourceType"`
	Description                string            `json:"description"`
	Label                      string            `json:"label"`
	Name                       string            `json:"name"`
	ResourceTypeID             string            `json:"resourceTypeId"`
	Storage                    []Storage         `json:"storage"`
	Units                      ResourceTypeUnits `json:"units,omitempty"`
}

type ResourceTypeUnits struct {
	Readings string `json:"readings"`
}

type Fields struct {
	Datatype  string `json:"datatype"`
	FieldName string `json:"fieldName"`
	Negative  bool   `json:"negative"`
}

type Storage struct {
	Fields   []Fields  `json:"fields"`
	Sampling string    `json:"sampling"`
	Start    time.Time `json:"start"`
	Type     string    `json:"type"`
}

// Devices is a slice of Device
type Devices []Device

// Devices represents a single device know to the API
type Device struct {
	Active           bool        `json:"active"`
	CreatedAt        time.Time   `json:"createdAt"`
	Description      string      `json:"description"`
	DeviceID         string      `json:"deviceId"`
	DeviceTypeID     string      `json:"deviceTypeId"`
	HardwareID       string      `json:"hardwareId"`
	HardwareIDNames  []string    `json:"hardwareIdNames"`
	HardwareIds      HardwareIds `json:"hardwareIds,omitempty"`
	OwnerID          string      `json:"ownerId"`
	ParentHardwareID []string    `json:"parentHardwareId"`
	Protocol         Protocol    `json:"protocol"`
	Tags             []string    `json:"tags"`
	UpdatedAt        time.Time   `json:"updatedAt"`
}

type Sensors struct {
	ProtocolID     string `json:"protocolId"`
	ResourceID     string `json:"resourceId"`
	ResourceTypeID string `json:"resourceTypeId"`
}

type Protocol struct {
	Protocol string    `json:"protocol"`
	Sensors  []Sensors `json:"sensors"`
}

type HardwareIds struct {
	EUI          string `json:"EUI,omitempty"`
	MAC          string `json:"MAC,omitempty"`
	MPAN         string `json:"MPAN,omitempty"`
	MPRN         string `json:"MPRN,omitempty"`
	SerialNumber string `json:"serialNumber,omitempty"`
}
