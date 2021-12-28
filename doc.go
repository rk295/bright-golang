/*
Package bright-golang is a client for the Bright App energy monitoring API.

After setting BRIGHT_USERNAME and BRIGHT_PASSWORD you could quickly see your
current electricity usage in Watts by doing:

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
*/
package bright
