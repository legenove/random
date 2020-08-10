package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strconv"
	"unsafe"
)

func ConcatenateStrings(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, i := range s {
		buffer.WriteString(i)
	}
	return buffer.String()
}

func ConcatenateInt64String(s ...int64) string {
	if len(s) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, i := range s {
		buffer.WriteString(strconv.FormatInt(int64(i), 10))
	}
	return buffer.String()
}

func Bytes2Str(v []byte) string {
	return *(*string)(unsafe.Pointer(&v))
}

func MergeTwoMapString(map1, map2 map[string]string) map[string]string {
	if len(map1) >= len(map2) {
		for k, v := range map2 {
			map1[k] = v
		}
		return map1
	} else {
		for k, v := range map1 {
			map2[k] = v
		}
		return map2
	}
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func CheckNormalKey(s string) bool {
	var rCharacter = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9\\-_]*$")
	return rCharacter.MatchString(s)
}

func IsInStringSlice(slice []string, e string) bool {
	for _, s := range slice {
		if e == s {
			return true
		}
	}
	return false
}