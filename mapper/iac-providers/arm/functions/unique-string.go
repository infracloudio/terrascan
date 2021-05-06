package functions

import (
	"github.com/google/uuid"
)

// UniqueString function returns a string UUID.
func UniqueString() string {
	return uuid.NewString()
}
