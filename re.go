package common

import "regexp"

const ()

func Regex(s, expr string) (string, error) {
	r, err := regexp.Compile(expr)
	if err != nil {
		return "", err
	}
	return r.FindString(s), nil
}
