package common

import (
	"log"
	"runtime/debug"
)

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic in goroutine：%v, stack: %v", err, string(debug.Stack()))
			}
		}()
		f()
	}()
}
