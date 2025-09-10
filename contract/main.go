package main

import (
	"contract-template/sdk"
	_ "contract-template/sdk" // ensure sdk is imported
	"encoding/json"
)

func main() {

}

// Convenience helper
func strptr(s string) *string { return &s }

// =============================
//go:wasmexported contract entrypoints
// =============================

//go:wasmexport ping
func Ping() *string {
	return strptr("pong")
}

//go:wasmexport echo
func Echo(msg *string) *string {
	if msg == nil {
		return strptr("")
	}
	sdk.Log("echo: " + *msg)
	return msg
}

// // --- State ---

// //go:wasmexport set_object
// func Set_object(key *string, value *string) *string {
// 	if key == nil || value == nil {
// 		return strptr("error: missing key/value")
// 	}
// 	sdk.StateSetObject(*key, *value)
// 	return strptr("ok")
// }

// //go:wasmexport get_object
// func Get_object(key *string) *string {
// 	if key == nil {
// 		return strptr("")
// 	}
// 	res := sdk.StateGetObject(*key)
// 	if res == nil {
// 		return strptr("")
// 	}
// 	return res
// }

// //go:wasmexport rm_object
// func Rm_object(key *string) *string {
// 	if key == nil {
// 		return strptr("error: missing key")
// 	}
// 	sdk.StateDeleteObject(*key)
// 	return strptr("ok")
// }

// // --- Env ---

//go:wasmexport get_env_json
func Get_env_json() *string {
	env := sdk.GetEnv()
	b, _ := json.Marshal(env)
	s := string(b)
	return &s
}

// //go:wasmexport get_env_key
// func Get_env_key(k *string) *string {
// 	if k == nil {
// 		return strptr("")
// 	}
// 	return sdk.GetEnvKey(*k)
// }

// // --- Balances ---

// //go:wasmexport get_balance
// func Get_balance(addr *string, asset *string) *string {
// 	if addr == nil || asset == nil {
// 		return strptr("error: missing addr/asset")
// 	}
// 	bal := sdk.GetBalance(sdk.Address(*addr), sdk.Asset(*asset))
// 	s := strconv.FormatInt(bal, 10)
// 	return &s
// }

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

// // --- Cross-contract ---

// // //wasmexport read_other
// // func read_other(contractID *string, key *string) *string {
// // 	if contractID == nil || key == nil {
// // 		return strptr("")
// // 	}
// // 	return sdk.ContractRead(contractID, key)
// // }

// // //wasmexport call_other
// // func call_other(contractID *string, method *string, payload *string, options *string) *string {
// // 	if contractID == nil || method == nil {
// // 		return strptr("error: missing contract/method")
// // 	}
// // 	return sdk.ContractCall(contractID, method, payload, options)
// // }
