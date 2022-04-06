package common

import (
	"crypto/md5"
	"fmt"
)

func Gen(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
