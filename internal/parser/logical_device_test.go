package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/d9ff/adaptec-raid-exporter/internal/command"
)

func TestScanner_ScanLogicalDeviceInformation(t *testing.T) {
	mock := command.MockRunner{BasePath: "testdata"}

	sc := New(mock)
	status, err := sc.ScanLogicalDeviceStatus()
	assert.NoError(t, err)
	assert.Equal(t, "Degraded", status)
}
