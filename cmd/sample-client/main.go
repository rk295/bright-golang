package main

import (
	"fmt"

	"github.com/rk295/bright-golang"
	"github.com/sirupsen/logrus"
)

func main() {

	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel) // Very chatty, including HTTP response body

	c, err := bright.NewClientFromEnv()
	if err != nil {
		panic(err)
	}
	c.WithLogger(logger)

	electricityW, err := c.GetElectricityCurrentWatts()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current electricity usage: %dW\n", electricityW)

}
