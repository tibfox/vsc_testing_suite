package contract_test

import (
	"embed"
	"encoding/json"
	"fmt"

	"testing"
	"vsc-node/lib/test_utils"
	"vsc-node/modules/db/vsc/contracts"

	ledgerDb "vsc-node/modules/db/vsc/ledger"
	stateEngine "vsc-node/modules/state-processing"

	"github.com/stretchr/testify/assert"
)

//go:embed artifacts/main.wasm
var ContractWasm []byte

const contractId = "vsctestcontract"

var _ = embed.FS{}

func TestContractA(t *testing.T) {
	ct := test_utils.NewContractTest()
	ct.RegisterContract(contractId, ContractWasm)
	ct.Deposit("hive:someone", 1000, ledgerDb.AssetHive) // deposit 1 HIVE
	ct.Deposit("hive:someone", 1000, ledgerDb.AssetHbd)  // deposit 1 HBD

	result, gasUsed, logs := ct.Call(stateEngine.TxVscCallContract{
		Self: stateEngine.TxSelf{
			TxId:                 "sometxid",
			BlockId:              "abcdef",
			Index:                69,
			OpIndex:              0,
			Timestamp:            "2025-09-03T00:00:00",
			RequiredAuths:        []string{"hive:someone"},
			RequiredPostingAuths: []string{},
		},
		ContractId: contractId,
		Action:     "show_transfer_allow",
		Payload:    json.RawMessage(""),
		RcLimit:    1000,
		Intents: []contracts.Intent{{
			Type: "transfer.allow",
			Args: map[string]string{
				"limit": "1.000",
				"token": "hive",
			},
		}},
	})
	assert.True(t, result.Success)                 // assert contract execution success
	assert.LessOrEqual(t, gasUsed, uint(10000000)) // assert this call uses no more than 10M WASM gas
	assert.GreaterOrEqual(t, len(logs), 1)         // assert at least 1 log emitted
	for i, logEntry := range logs {
		fmt.Printf("Log %d: %v\n", i, logEntry)
	}
}

func TestContractB(t *testing.T) {
	ct := test_utils.NewContractTest()
	ct.RegisterContract(contractId, ContractWasm)
	ct.Deposit("hive:someone", 1000, ledgerDb.AssetHive) // deposit 1 HIVE
	ct.Deposit("hive:someone", 1000, ledgerDb.AssetHbd)  // deposit 1 HBD

	result, gasUsed, logs := ct.Call(stateEngine.TxVscCallContract{
		Self: stateEngine.TxSelf{
			TxId:                 "sometxid",
			BlockId:              "abcdef",
			Index:                69,
			OpIndex:              0,
			Timestamp:            "2025-09-03T00:00:00",
			RequiredAuths:        []string{"hive:someone"},
			RequiredPostingAuths: []string{},
		},
		ContractId: contractId,
		Action:     "show_transfer_allow",
		Payload:    json.RawMessage(""),
		RcLimit:    1000,
		Intents: []contracts.Intent{{
			Type: "transfer.allow",
			Args: map[string]string{
				"limit": "1.000",
				"token": "hive",
			},
		}},
	})
	assert.True(t, result.Success)                 // assert contract execution success
	assert.LessOrEqual(t, gasUsed, uint(10000000)) // assert this call uses no more than 10M WASM gas
	assert.GreaterOrEqual(t, len(logs), 1)         // assert at least 1 log emitted
	for i, logEntry := range logs {
		fmt.Printf("Log %d: %v\n", i, logEntry)
	}
}
