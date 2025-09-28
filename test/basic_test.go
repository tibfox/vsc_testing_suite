package contract_test

import (
	"testing"

	"vsc-node/modules/db/vsc/contracts"
	ledgerDb "vsc-node/modules/db/vsc/ledger"
)

// Simple individual tests
func TestShowTA(t *testing.T) {
	CallContract(t, "show_transfer_allow", PayloadToJSON(""), []contracts.Intent{
		{Type: "transfer.allow", Args: map[string]string{"limit": "1.000", "token": "hive"}},
	}, true)
}

func TestEcho(t *testing.T) {
	CallContract(t, "echo", PayloadToJSON("test"), nil, true)
}

func TestSomeLogs(t *testing.T) {
	CallContract(t, "emit_event_logs", PayloadToJSON(""), nil, true)
}

func TestPing(t *testing.T) {
	CallContract(t, "ping", PayloadToJSON(""), nil, true)
}

// Table-driven example
func TestContractActions(t *testing.T) {
	tests := []ContractTestCase{
		{"ShowTA", "show_transfer_allow", "", []contracts.Intent{{Type: "transfer.allow", Args: map[string]string{"limit": "1.000", "token": "hive"}}}, true},
		{"Echo", "echo", "test", nil, true},
		{"Ping", "ping", "", nil, true},
		{"Show Balance", "get_balance", map[string]string{"address": "hive:someone", "asset": "hbd"}, nil, false},
	}

	RunContractTests(t, tests)
}

func TestContractStateChanges(t *testing.T) {
	tests := []ContractTestCase{
		{"Set First", "set_object", map[string]string{"stateKey": "testkey", "stateValue": "testval"}, nil, true},
		{"Get First", "get_object", "testkey", nil, true},
		{"Set Second", "set_object", map[string]string{"stateKey": "testkey", "stateValue": "testval2"}, nil, true},
		{"Get Second", "get_object", "testkey", nil, true},
		{"Delete", "rm_object", "testkey", nil, true},
		{"Get Third", "get_object", "testkey", nil, true},
	}
	RunContractTests(t, tests)
}

func TestContractGetBalances(t *testing.T) {
	testBeforeDeposite := []ContractTestCase{
		{"Get Balance Hive", "get_balance", map[string]string{"address": "hive:someone", "asset": "hive"}, nil, true},
		{"Get Balance HBD", "get_balance", map[string]string{"address": "hive:someone", "asset": "hive"}, nil, true},
	}
	RunContractTests(t, testBeforeDeposite)
	SharedCT.Deposit("hive:someone", 1000, ledgerDb.AssetHive)
	SharedCT.Deposit("hive:someone", 1000, ledgerDb.AssetHbd)

	testAfterDeposite := []ContractTestCase{
		{"Get Balance Hive", "get_balance", map[string]string{"address": "hive:someone", "asset": "hive"}, nil, true},
		{"Get Balance HBD", "get_balance", map[string]string{"address": "hive:someone", "asset": "hive"}, nil, true},
	}
	RunContractTests(t, testAfterDeposite)
}
