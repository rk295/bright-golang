# Go client library for the Bright API

[![Go Reference](https://pkg.go.dev/badge/github.com/rk295/bright-golang.svg)](https://pkg.go.dev/github.com/rk295/bright-golang)
[![CI](https://github.com/rk295/bright-golang/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/rk295/bright-golang/actions/workflows/go.yml)

Package bright-golang is a client for the Bright App energy monitoring API.

Not all the API endpoints are implemented, however there should be sufficient
to at least get Electricity and Gas usage out. There are a few helper functions
which attempt to figure out which resource holds the current usage (in Watts) for both
gas and electricity, they are the simplest ways of using the package.

A username and password for the API is required, you can either pass those into
NewClient() or you can have them read from the environment by calling
NewClientFromEnv().  This uses the BRIGHT_USERNAME and BRIGHT_PASSWORD
environment variables.

To quickly see your current electricity usage in Watts you could do somehting like:

	package main
	import (
		"fmt"
		"github.com/rk295/bright-golang"
	)

	func main() {
		c, _ := bright.NewClientFromEnv()
		electricityWatts, _ := c.GetElectricityCurrentWatts()
		fmt.Println(electricityWatts)
	}

Ouput:
	566

Indicating that 566W of electricity is being used now.

## Sample apps

### Simple CLI

In [cmd/sample-client](cmd/sample-client) there is a simple sample application. Have a look at the [README.md](cmd/sample-client/README.md) in there for more.

### MQTT

In [cmd/sample-mqtt](cmd/sample-mqtt) there is a MQTT client which will connect to Bright via MQTT if they have enabled it for you. There is a [README.md](cmd/sample-mqtt/README.md) in there with more info.

### Prometheus Exporter

in [cmd/bright-exporter-api](cmd/bright-exporter-api) there is a Prometheus exporter which fetches data from the Bright API. Again have a look at the [README.md](cmd/bright-exporter-api/README.md) in that directory.
