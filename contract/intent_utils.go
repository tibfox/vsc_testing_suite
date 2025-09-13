package main

import (
	"strconv"
	"vsc_testing_suite/sdk"
)

// New struct for transfer.allow args
type TransferAllow struct {
	Limit float64
	Token sdk.Asset
}

// TODO: define your accepted tokens here
var validAssets = []string{sdk.AssetHbd.String(), sdk.AssetHive.String()}

// Helper function to validate token
func isValidAsset(token string) bool {
	for _, a := range validAssets {
		if token == a {
			return true
		}
	}
	return false
}

// Helper function to get the first transfer.allow intent (if exists)
func GetFirstTransferAllow(intents []sdk.Intent) *TransferAllow {
	for _, intent := range intents {
		if intent.Type == "transfer.allow" {
			token := intent.Args["token"]
			// if we have an transfer.allow intent but the asset is not valid
			if !isValidAsset(token) {
				sdk.Abort("invalid intent token")
			}
			limitStr := intent.Args["limit"]
			limit, err := strconv.ParseFloat(limitStr, 64)
			if err != nil {
				sdk.Abort("invalid intent limit")
			}
			ta := &TransferAllow{
				Limit: limit,
				Token: sdk.Asset(token),
			}
			return ta
		}
	}
	return nil
}
