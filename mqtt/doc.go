package mqtt

/*
Package bright-golang/mqtt is a set of types which aide in the Unmarshalling
of the JSON messages received via MQTT.

Given a MQTT message payload in m the following will print the current
Electricity usage in Watts

	if err := json.Unmarshal(m.Payload(), &decoded); err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(decoded.ElecMtr.Metering.HistoricalConsumption.InstantaneousDemand)

*/
