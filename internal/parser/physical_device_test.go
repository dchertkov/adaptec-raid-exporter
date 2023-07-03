package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/d9ff/adaptec-raid-exporter/internal/command"
)

func TestScanner_ScanPhysicalDevices(t *testing.T) {
	mock := command.MockRunner{BasePath: "testdata"}

	sc := New(mock)
	devices, err := sc.ScanPhysicalDevices()
	assert.NoError(t, err)
	assert.Len(t, devices, 15)
	assert.Equal(t, PhysicalDevice{
		IsHardDrive:   true,
		Device:        "0",
		SerialNumber:  "xxx",
		State:         "Online",
		SMARTWarnings: 7,
	}, devices[0])
	assert.Equal(t, PhysicalDevice{
		IsHardDrive:  true,
		Device:       "1",
		SerialNumber: "yyy",
		State:        "Online",
	}, devices[1])
}
