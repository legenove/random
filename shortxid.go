package random

import "github.com/rs/xid"


func Xid() string {
	// unique short id
	guid := xid.New()
	return guid.String()
}