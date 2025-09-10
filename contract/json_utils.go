package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"vsc_testing_suite/sdk"
)

// Conversions from/to json strings

func ToJSON[T any](v T, objectType string) string {
	b, err := json.Marshal(v)
	if err != nil {
		sdk.Abort(fmt.Sprintf("failed to marshal %s\nInput data:%+v\nError: %v:", objectType, v, err))
	}
	return string(b)
}

func FromJSON[T any](data string, objectType string) *T {
	data = strings.TrimSpace(data)
	var v T
	if err := json.Unmarshal([]byte(data), &v); err != nil {
		sdk.Abort(fmt.Sprintf(
			"failed to unmarshal %s\nInput data:%s\nError: %v:", objectType, data, err))
	}
	return &v
}
