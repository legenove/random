package random

import (
	"fmt"
	"hash/crc32"
	"strconv"
)

func GetShardString(d interface{}, n uint32) string {
	var s string
	switch v := d.(type) {
	case int:
		s = strconv.Itoa(v)
	case string:
		s = v
	}
	c := crc32.ChecksumIEEE([]byte(s))
	return fmt.Sprintf("%x", c%n)
}


func GetShardNum(d interface{}, n uint32) int {
	var s string
	switch v := d.(type) {
	case int:
		s = strconv.Itoa(v)
	case string:
		s = v
	}
	c := crc32.ChecksumIEEE([]byte(s))
	return int(c%n)
}

