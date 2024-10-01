package app

import (
	"testing"

	dbm "github.com/cosmos/cosmos-db"

	"cosmossdk.io/log"

	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/reece/myproject/app/params"
)

// MakeEncodingConfig creates a new EncodingConfig with all modules registered. For testing only
func MakeEncodingConfig(t testing.TB) params.EncodingConfig {
	t.Helper()
	// we "pre"-instantiate the application for getting the injected/configured encoding configuration
	// note, this is not necessary when using app wiring, as depinject can be directly used (see root_v2.go)
	tempApp := NewChainApp(
		log.NewNopLogger(),
		dbm.NewMemDB(),
		nil,
		true,
		simtestutil.NewAppOptionsWithFlagHome(t.TempDir()),
		[]wasmkeeper.Option{},
	)
	return makeEncodingConfig(tempApp)
}

func makeEncodingConfig(tempApp *ChainApp) params.EncodingConfig {
	encodingConfig := params.EncodingConfig{
		InterfaceRegistry: tempApp.InterfaceRegistry(),
		Codec:             tempApp.AppCodec(),
		TxConfig:          tempApp.TxConfig(),
		Amino:             tempApp.LegacyAmino(),
	}
	return encodingConfig
}
