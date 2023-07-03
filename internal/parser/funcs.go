package parser

import (
	"bytes"
	"strconv"
)

func parseString(line []byte, name []byte) (value string, ok bool) {
	index := bytes.Index(line, name)
	if index == -1 {
		return
	}

	ok = true
	value = string(bytes.Trim(line[index+len(name):], " :"))
	return
}

func parseInt(line []byte, name []byte) (value int, ok bool) {
	var s string
	s, ok = parseString(line, name)
	if !ok {
		return
	}
	value, _ = strconv.Atoi(s)
	return
}
