package parser

import (
	"bufio"
	"bytes"
)

var (
	logicalDeviceStatus = []byte("Status of Logical Device")
)

func (sc *Scanner) ScanLogicalDeviceStatus() (string, error) {
	output, err := sc.command.Exec("arcconf", "getconfig", "1", "ld")
	if err != nil {
		return "", err
	}

	scan := bufio.NewScanner(bytes.NewReader(output))
	for scan.Scan() {
		line := scan.Bytes()

		if s, ok := parseString(line, logicalDeviceStatus); ok {
			return s, nil
		}
	}

	return "", nil
}
