package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/d9ff/adaptec-raid-exporter/internal/command"
	"github.com/d9ff/adaptec-raid-exporter/internal/parser"
)

var (
	logicalDeviceStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "adaptec_logical_device_status",
		Help: "Status of Logical Devices",
	}, []string{"status"})

	physicalDeviceState = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "adaptec_physical_device_state",
		Help: "State of physical devices",
	}, []string{"device", "state"})

	physicalDeviceSMARTWarnings = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "adaptec_physical_device_smart_warnings",
		Help: "SMART warnings of physical devices",
	}, []string{"device"})
)

func init() {
	prometheus.MustRegister(logicalDeviceStatus)
	prometheus.MustRegister(physicalDeviceState)
	prometheus.MustRegister(physicalDeviceSMARTWarnings)
}

func main() {
	addr := flag.String("web.listen-address", ":9101", "listen address")
	flag.Parse()

	cmd := &command.ExecRunner{}
	p := parser.New(cmd)

	go func() {
		t := time.NewTicker(time.Second * 30)

		for range t.C {
			status, err := p.ScanLogicalDeviceStatus()
			if err != nil {
				panic(err)
			}

			logicalDeviceStatus.Reset()
			logicalDeviceStatus.WithLabelValues(status).Set(1)

			pd, err := p.ScanPhysicalDevices()
			if err != nil {
				panic(err)
			}

			physicalDeviceSMARTWarnings.Reset()
			physicalDeviceState.Reset()
			for _, device := range pd {
				physicalDeviceSMARTWarnings.WithLabelValues(device.Device).Set(float64(device.SMARTWarnings))
				physicalDeviceState.WithLabelValues(device.Device, device.State).Set(1)
			}
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(*addr, nil); err != nil {
		panic(err)
	}
}
