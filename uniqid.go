package random

import (
	"github.com/legenove/utils"
	"math"
	"strconv"
	"time"
)

// Uniqid 随机数，length是需要返回的长度，只支持12~18位
// 20年不会有问题
func Uniqid(length int) int64 {
	if length < 12 || length > 18 {
		return 0
	}
	// 0.1秒级别
	prefix := time.Now().UnixNano()/100000000 - StartUnixTime*10
	cut := int64(math.Pow10(length - 10))
	suffix := <-ch % cut
	return prefix*cut + suffix
}

// UniqStringId 随机数字符串，只支持19~32位
// 		length是需要返回的长度，
// 		start支持1000以内的数字，默认为1
func UniqStringId(length int, starts ...int) string {
	if length < 20 || length > 32 {
		return ""
	}
	var start = 1
	if len(starts) > 0 {
		start = starts[0]
	}
	var sLength = 3
	if start > 1000 {
		return ""
	} else if start < 10 {
		sLength = 1
	} else if start < 100 {
		sLength = 2
	}
	// 0.01毫秒级别
	prefix := strconv.FormatInt(time.Now().UnixNano()/10000 - StartUnixTime*100000, 10)
	cutLen := 2 + length - 20
	cut := int64(math.Pow10(cutLen))
	suffix := strconv.FormatInt(<-ch % cut, 10)
	padLength := length - len(prefix) - sLength - len(suffix)
	ss := make([]string, padLength+3)
	prefixIndex := padLength+1-(cutLen-len(suffix))
	ss[0] = strconv.Itoa(start)
	ss[prefixIndex] = prefix
	ss[padLength+2] = suffix
	for i := 1; i < prefixIndex; i ++ {
		if ss[i] == "" {
			ss[i] = "0"
		}
	}
	for i := prefixIndex; i < padLength+2; i ++ {
		if ss[i] == "" {
			ss[i] = "0"
		}
	}
	return utils.ConcatenateStrings(ss...)

}
