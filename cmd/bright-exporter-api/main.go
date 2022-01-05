package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"

	"github.com/rk295/bright-golang"
)

type Data struct {
	client *bright.Client
	logger *logrus.Logger
}

const (
	exporterPortEnv     = "PORT"
	exporterDefaultPort = "9998"
)

var (
	electricityDetails = prometheus.NewDesc(
		"bright_energy",
		"electricity (W))",
		[]string{"kind"}, nil,
	)
)

func main() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	exporterPort := os.Getenv(exporterPortEnv)
	if exporterPort == "" {
		logger.Debugf("%s not set, using default port of %s", exporterPortEnv, exporterDefaultPort)
		exporterPort = exporterDefaultPort
	}

	c, err := bright.NewClientFromEnv()
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	c.WithLogger(logger).WithLevel(logger.GetLevel())

	exporter := &Data{
		client: c,
		logger: logger,
	}

	prometheus.MustRegister(exporter)
	http.Handle("/metrics", promhttp.Handler())
	logger.Infof("starting metrics server on port %s", exporterPort)
	err = http.ListenAndServe(fmt.Sprintf(":%s", exporterPort), nil)
	if err != nil {
		logger.Error(err)
		logger.Exit(1)
	}

}

func (d Data) Describe(ch chan<- *prometheus.Desc) {
	ch <- electricityDetails
}

func (d Data) Collect(ch chan<- prometheus.Metric) {
	electricityCurrent, err := d.client.GetElectricityCurrentWatts()
	if err != nil {
		d.logger.Error(err)
		return
	}
	ch <- prometheus.MustNewConstMetric(
		electricityDetails,
		prometheus.GaugeValue,
		float64(electricityCurrent),
		[]string{"electricity"}...,
	)

	gasCurrent, err := d.client.GetGasCurrentWatts()
	if err != nil {
		d.logger.Error(err)
		return
	}
	ch <- prometheus.MustNewConstMetric(
		electricityDetails,
		prometheus.GaugeValue,
		float64(gasCurrent),
		[]string{"gas"}...,
	)
}
