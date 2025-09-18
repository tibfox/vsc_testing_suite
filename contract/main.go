package main

import (
	"fmt"
	"strconv"
	"vsc_testing_suite/sdk"
	_ "vsc_testing_suite/sdk" // ensure sdk is imported
)

func main() {

}

// Convenience helper
func strptr(s string) *string { return &s }

// =============================
//go:wasmexported contract entrypoints
// =============================

//go:wasmexport ping
func Ping(none *string) *string {
	sdk.Log("pong")
	return strptr("pong")
}

//go:wasmexport echo
func Echo(msg *string) *string {
	if msg == nil {
		return strptr("")
	}
	sdk.Log(*msg)
	return msg
}

// // --- State ---

type SetObjectPayloadArgs struct {
	Key   string `json:"stateKey"`
	Value string `json:"stateValue"`
}

//go:wasmexport set_object
func Set_object(payload *string) *string {
	input := FromJSON[SetObjectPayloadArgs](*payload, "set_object arguments")
	if input.Key == "" || input.Value == "" {
		return strptr("error: empty stateKey or stateValue")
	}
	sdk.StateSetObject(input.Key, input.Value)
	sdk.Log("ok")
	return strptr("ok")
}

//go:wasmexport get_object
func Get_object(stateKey *string) *string {
	if stateKey == nil {
		return strptr("")
	}
	res := sdk.StateGetObject(*stateKey)
	if res == nil || *res == "" {
		sdk.Log("stateKey not found in contract state")
		return strptr("stateKey not found in contract state")
	}
	sdk.Log(*res)
	return res
}

//go:wasmexport rm_object
func Rm_object(stateKey *string) *string {
	if stateKey == nil {
		return strptr("error: missing stateKey")
	}
	sdk.StateDeleteObject(*stateKey)
	sdk.Log("ok")
	return strptr("ok")
}

// --- Env ---

//go:wasmexport get_env_json
func GetEnvJSON(none *string) *string {
	env := sdk.GetEnv()
	j := ToJSON(env, "env")
	sdk.Log(j)
	return strptr(j)
}

//go:wasmexport get_env_key
func Get_env_key(k *string) *string {
	if k == nil {
		return strptr("env key not found")
	}
	envVal := sdk.GetEnvKey(*k)
	sdk.Log(*envVal)
	return envVal
}

// --- Intent checks ---

//go:wasmexport show_transfer_allow
func ShowIntent(none *string) *string {
	env := sdk.GetEnv()
	transferAllowIntent := GetFirstTransferAllow(env.Intents)
	if transferAllowIntent == nil {
		return strptr("none")
	}
	r := fmt.Sprintf("%f %s", transferAllowIntent.Limit, transferAllowIntent.Token.String())
	sdk.Log(r)
	return strptr(r)
}

// --- Balances ---

type showBalanceArgs struct {
	Address sdk.Address `json:"address"`
	Asset   sdk.Asset   `json:"asset"`
}

//go:wasmexport get_balance
func Get_balance(payload *string) *string {
	input := FromJSON[showBalanceArgs](*payload, "get_balance arguments")
	bal := sdk.GetBalance(input.Address, input.Asset)
	s := strconv.FormatInt(bal, 10)
	sdk.Log(s)
	return strptr(s)
}

// // --- Token flows ---

// //go:wasmexport draw
// func Draw(amount *string, asset *string) *string {
// 	if amount == nil || asset == nil {
// 		return strptr("error: missing amount/asset")
// 	}
// 	amt, err := strconv.ParseInt(*amount, 10, 64)
// 	if err != nil {
// 		return strptr("error: bad amount")
// 	}
// 	env := sdk.GetEnv()
// 	transferAllow := GetFirstTransferAllow(env.Intents)
// 	if transferAllow.Limit < amt{
// 		strptr("intent too low")
// 	}
// 	if transferAllow.Token != sdk.Asset(*asset){
// 		strptr("intent token not equal to asset")
// 	}
// 	sdk.HiveDraw(amt, sdk.Asset(*asset))
// 	return strptr("ok")
// }

// //go:wasmexport transfer
// func Transfer(to *string, amount *string, asset *string) *string {
// 	if to == nil || amount == nil || asset == nil {
// 		return strptr("error: missing args")
// 	}
// 	amt, err := strconv.ParseInt(*amount, 10, 64)
// 	if err != nil {
// 		return strptr("error: bad amount")
// 	}
// 	sdk.HiveTransfer(sdk.Address(*to), amt, sdk.Asset(*asset))
// 	return strptr("ok")
// }

// //go:wasmexport withdraw
// func Withdraw(to *string, amount *string, asset *string) *string {
// 	if to == nil || amount == nil || asset == nil {
// 		return strptr("error: missing args")
// 	}
// 	amt, err := strconv.ParseInt(*amount, 10, 64)
// 	if err != nil {
// 		return strptr("error: bad amount")
// 	}
// 	sdk.HiveWithdraw(sdk.Address(*to), amt, sdk.Asset(*asset))
// 	return strptr("ok")
// }

// --- Cross-contract ---

// //wasmexport read_other
// func read_other(contractID *string, key *string) *string {
// 	if contractID == nil || key == nil {
// 		return strptr("")
// 	}
// 	return sdk.ContractRead(contractID, key)
// }

// //wasmexport call_other
// func call_other(contractID *string, method *string, payload *string, options *string) *string {
// 	if contractID == nil || method == nil {
// 		return strptr("error: missing contract/method")
// 	}
// 	return sdk.ContractCall(contractID, method, payload, options)
// }
