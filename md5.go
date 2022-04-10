package common

import (
	"crypto/md5"
	"fmt"
)

func Md5(b []byte) string {
	has := md5.Sum(b)
	return fmt.Sprintf("%x", has)
}
