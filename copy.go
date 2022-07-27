package common

import "github.com/jinzhu/copier"

func CopyWithOption(toValue interface{}, fromValue interface{}) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{IgnoreEmpty: true, DeepCopy: true})
}
