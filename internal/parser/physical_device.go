package parser

import (
	"bufio"
	"bytes"
)

var (
	deviceNumber        = []byte("Device #")
	deviceHardDrive     = []byte("Device is a Hard drive")
	deviceState         = []byte("  State")
	deviceSMARTWarnings = []byte("S.M.A.R.T. warnings")
	deviceSerialNumber  = []byte("Serial number")
)

type PhysicalDevice struct {
	IsHardDrive   bool
	Device        string
	SerialNumber  string
	State         string
	SMARTWarnings int
}

func (sc *Scanner) ScanPhysicalDevices() (devices []PhysicalDevice, err error) {
	output, err := sc.command.Exec("arcconf", "getconfig", "1", "pd")

	scan := bufio.NewScanner(bytes.NewReader(output))

	var i int
	var s string
	var n int
	var ok bool

	for scan.Scan() {
		line := scan.Bytes()

		index := bytes.Index(line, deviceNumber)
		if index != -1 {
			devices = append(devices, PhysicalDevice{
				Device: string(line[index+len(deviceNumber):]),
			})
			i = len(devices) - 1
			continue
		}

		if bytes.Contains(line, deviceHardDrive) {
			devices[i].IsHardDrive = true
			continue
		}

		if s, ok = parseString(line, deviceState); ok {
			devices[i].State = s
			continue
		}

		if n, ok = parseInt(line, deviceSMARTWarnings); ok {
			devices[i].SMARTWarnings = n
			continue
		}

		if s, ok = parseString(line, deviceSerialNumber); ok {
			devices[i].SerialNumber = s
			continue
		}
	}

	filtered := devices[:0]
	for i = range devices {
		if !devices[i].IsHardDrive {
			continue
		}
		filtered = append(filtered, devices[i])
	}

	return filtered, nil
}
