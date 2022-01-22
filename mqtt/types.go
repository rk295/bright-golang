package mqtt

type Power struct {
	ElecMtr   Meter  `json:"elecMtr,omitempty"`
	GasMtr    Meter  `json:"gasMtr,omitempty"`
	Ts        string `json:"ts,omitempty"`
	Hversion  string `json:"hversion,omitempty"`
	Time      string `json:"time,omitempty"`
	ZbSoftVer string `json:"zbSoftVer,omitempty"`
	Gmtime    int    `json:"gmtime,omitempty"`
	Pan       Pan    `json:"pan,omitempty"`
	SmetsVer  string `json:"smetsVer,omitempty"`
	Ets       string `json:"ets,omitempty"`
	Gid       string `json:"gid,omitempty"`
}

type Pan struct {
	Rssi   string `json:"rssi,omitempty"`
	Status string `json:"status,omitempty"`
	NPAN   string `json:"nPAN,omitempty"`
	Join   string `json:"join,omitempty"`
	Lqi    string `json:"lqi,omitempty"`
}

type Meter struct {
	// Price            Price            `json:"0700,omitempty"`
	Metering         Metering         `json:"0702,omitempty"`
	Prepayment       Prepayment       `json:"0705,omitempty"`
	DeviceManagement DeviceManagement `json:"0708,omitempty"`
}

// 0702
type Metering struct {
	ReadingInformationSet ReadingInformationSet `json:"00,omitempty"`
	MeterStatus           MeterStatus           `json:"02,omitempty"`
	Formatting            Formatting            `json:"03,omitempty"`
	HistoricalConsumption HistoricalConsumption `json:"04,omitempty"`
	ZeroC                 ZeroC                 `json:"0C,omitempty"`
}

// 0702.00
type ReadingInformationSet struct {
	CurrentSummationDelivered ZBInt  `json:"00,omitempty"`
	CurrentSummationReceived  string `json:"01,omitempty"`
	CurrentMaxDemandDelivered string `json:"02,omitempty"`
	ReadingSnapshotTime       string `json:"07,omitempty"`
	SupplyStatus              string `json:"14,omitempty"`
}

// 0702.02
type MeterStatus struct {
	MeterStatus string `json:"00,omitempty"`
}

// 0702.03
type Formatting struct {
	UnitofMeasurestring      string `json:"00,omitempty"`
	Multiplier               ZBInt  `json:"01,omitempty"`
	Divisor                  ZBInt  `json:"02,omitempty"`
	SummationFormatting      string `json:"03,omitempty"`
	DemandFormatting         string `json:"04,omitempty"`
	MeteringDeviceType       string `json:"06,omitempty"`
	SiteID                   string `json:"07,omitempty"`
	MeterSerialNumber        string `json:"08,omitempty"`
	AlternativeUnitofMeasure string `json:"12,omitempty"`
}

// 0702.04
type HistoricalConsumption struct {
	InstantaneousDemand              ZBInt `json:"00,omitempty"`
	CurrentDayConsumptionDelivered   ZBInt `json:"01,omitempty"`
	CurrentWeekConsumptionDelivered  ZBInt `json:"30,omitempty"`
	CurrentMonthConsumptionDelivered ZBInt `json:"40,omitempty"`
}

// 0702.0C
type ZeroC struct {
	CurrentDayConsumptionDelivered   ZBInt `json:"01,omitempty"`
	CurrentWeekConsumptionDelivered  ZBInt `json:"30,omitempty"`
	CurrentMonthConsumptionDelivered ZBInt `json:"40,omitempty"`
}

// 0705
type Prepayment struct {
	PrepaymentInformationSet PrepaymentInformationSet `json:"00,omitempty"`
}

// 0705.00
type PrepaymentInformationSet struct {
	PaymentControlConfiguration string `json:"00,omitempty"`
	CreditRemaining             string `json:"01,omitempty"`
}

// 0708
type DeviceManagement struct {
	SupplierControlAttributeSet SupplierControlAttributeSet `json:"01,omitempty"`
}

// 0708.01
type SupplierControlAttributeSet struct {
	ProviderName string `json:"01,omitempty"`
}
