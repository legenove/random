package random

import (
	"github.com/legenove/utils"
	"github.com/nu7hatch/gouuid"
	"strconv"
	"strings"
	"time"
)

func UuidV4() string {
	// 产生的uuid直接去掉"-"
	u, err := uuid.NewV4()
	if err != nil {
		return utils.GetMD5Hash(strconv.FormatInt(time.Now().UnixNano(), 10))
	}
	return strings.Replace(u.String(), "-", "", 4)
}

func UuidV5() string {
	idv4, err := uuid.NewV4()
	if err != nil {
		return utils.GetMD5Hash(strconv.FormatInt(time.Now().UnixNano(), 10))
	}
	rands := utils.IntToByte(<-ch)
	id, err := uuid.NewV5(idv4, rands)
	if err != nil {
		return utils.GetMD5Hash(strconv.FormatInt(time.Now().UnixNano(), 10))
	}
	return strings.Replace(id.String(), "-", "", 4)
}
