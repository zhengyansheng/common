package common

import (
	"fmt"

	"github.com/kr/pretty"
)

func Indent(x interface{}) {
	fmt.Printf("%# v", pretty.Formatter(x))
}
