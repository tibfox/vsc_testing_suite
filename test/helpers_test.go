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

var _ = embed.FS{} // just so "embed" can be imported

// Shared contract state across tests
var SharedCT *test_utils.ContractTest

const ContractID = "vsctestcontract"

//go:embed artifacts/main.wasm
var ContractWasm []byte

// SetupShared initializes the shared contract state only once
func SetupShared() *test_utils.ContractTest {
	if SharedCT == nil {
		ct := test_utils.NewContractTest()
		SharedCT = &ct
		SharedCT.RegisterContract(ContractID, ContractWasm)
		SharedCT.Deposit("hive:someone", 1000, ledgerDb.AssetHive)
		SharedCT.Deposit("hive:someone", 1000, ledgerDb.AssetHbd)
	}
	return SharedCT
}

// CallContract executes a contract action and asserts basic success
func CallContract(t *testing.T, action string, payload json.RawMessage, intents []contracts.Intent, expectLogs bool) (stateEngine.TxResult, uint, []string) {
	ct := SetupShared()
	result, gasUsed, logs := ct.Call(stateEngine.TxVscCallContract{
		Self: stateEngine.TxSelf{
			TxId:                 fmt.Sprintf("%s-tx", action),
			BlockId:              "block1",
			Index:                0,
			OpIndex:              0,
			Timestamp:            "2025-09-03T00:00:00",
			RequiredAuths:        []string{"hive:someone"},
			RequiredPostingAuths: []string{},
		},
		ContractId: ContractID,
		Action:     action,
		Payload:    payload,
		RcLimit:    1000,
		Intents:    intents,
	})

	PrintLogs(logs)
	PrintErrorIfFailed(result)

	assert.True(t, result.Success, "Contract action failed")
	assert.LessOrEqual(t, gasUsed, uint(10_000_000), "Gas exceeded limit")
	if expectLogs {
		assert.GreaterOrEqual(t, len(logs), 1, "Expected at least 1 log")
	}
	return result, gasUsed, logs
}

// PrintLogs prints all logs from a contract call
func PrintLogs(logs []string) {
	for i, logEntry := range logs {
		fmt.Printf("Log %d: %v\n", i, logEntry)
	}
}

// PrintErrorIfFailed prints error if the contract call failed
func PrintErrorIfFailed(result stateEngine.TxResult) {
	if !result.Success {
		fmt.Println(result.Err)
	}
}

// ToJSONRaw converts Go objects to json.RawMessage
func ToJSONRaw(v any) json.RawMessage {
	b, err := json.Marshal(v)
	if err != nil {
		panic(fmt.Sprintf("failed to marshal JSON: %v", err))
	}
	return b
}

// PayloadToJSON safely converts payloads to json.RawMessage
func PayloadToJSON(v any) json.RawMessage {
	switch val := v.(type) {
	case string:
		return json.RawMessage(fmt.Sprintf("%q", val)) // add quotes for strings
	case json.RawMessage:
		return val
	default:
		return ToJSONRaw(val)
	}
}

type ContractTestCase struct {
	Name       string
	Action     string
	Payload    any
	Intents    []contracts.Intent
	ExpectLogs bool
}

func RunContractTests(t *testing.T, tests []ContractTestCase) {
	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.Name, func(t *testing.T) {
			CallContract(t, tt.Action, PayloadToJSON(tt.Payload), tt.Intents, tt.ExpectLogs)
		})
	}
}
