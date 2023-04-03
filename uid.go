package common

import (
	"math/rand"
	"time"
	
	"github.com/google/uuid"
)

func NewUUID() (string, error) {
	rand.Seed(time.Now().UnixNano())
	u, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
