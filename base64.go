package common

import "encoding/base64"

func EncodeToString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
