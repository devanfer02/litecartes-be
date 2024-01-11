package utils

import (
	"strings"

	"github.com/google/uuid"
)

const (
    prefix = "LTC-APP-"
)

func ValidUUID(s string) bool {
    s = strings.Replace(s, prefix, "", 1)

    if err := uuid.Validate(s); err != nil {
        return false 
    }

    return true 
}

func CreateUUID() string {
    s := uuid.NewString()
    s = prefix + s 

    return s 
}