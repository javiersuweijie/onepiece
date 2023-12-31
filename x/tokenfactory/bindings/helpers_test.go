package bindings_test

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/crypto/ed25519"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/javiersuweijie/onepiece/app"

	"github.com/CosmWasm/wasmd/x/wasm"
	"github.com/CosmWasm/wasmd/x/wasm/keeper"
)

func CreateTestInput() (*app.App, sdk.Context) {
	var emptyWasmOpts []wasm.Option

	osmosis := app.Setup(&testing.T{}, emptyWasmOpts...)
	ctx := osmosis.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "osmosis-1", Time: time.Now().UTC()})
	return osmosis, ctx
}

func FundAccount(t *testing.T, ctx sdk.Context, osmosis *app.App, acct sdk.AccAddress) {
	t.Helper()
	// TODO:
	// err := simapp.FundAccount(osmosis.BankKeeper, ctx, acct, sdk.NewCoins(
	// 	sdk.NewCoin("uosmo", sdk.NewInt(10000000000)),
	// ))
	// require.NoError(t, err)
}

// we need to make this deterministic (same every test run), as content might affect gas costs
func keyPubAddr() (crypto.PrivKey, crypto.PubKey, sdk.AccAddress) {
	key := ed25519.GenPrivKey()
	pub := key.PubKey()
	addr := sdk.AccAddress(pub.Address())
	return key, pub, addr
}

func RandomAccountAddress() sdk.AccAddress {
	_, _, addr := keyPubAddr()
	return addr
}

func RandomBech32AccountAddress() string {
	return RandomAccountAddress().String()
}

func storeReflectCode(t *testing.T, ctx sdk.Context, tokenz *app.App, addr sdk.AccAddress) uint64 {
	t.Helper()
	wasmCode, err := os.ReadFile("./testdata/token_reflect.wasm")
	require.NoError(t, err)

	contractKeeper := keeper.NewDefaultPermissionKeeper(tokenz.WasmKeeper)
	codeID, _, err := contractKeeper.Create(ctx, addr, wasmCode, nil)
	require.NoError(t, err)

	return codeID
}

func instantiateReflectContract(t *testing.T, ctx sdk.Context, tokenz *app.App, funder sdk.AccAddress) sdk.AccAddress {
	t.Helper()
	initMsgBz := []byte("{}")
	contractKeeper := keeper.NewDefaultPermissionKeeper(tokenz.WasmKeeper)
	codeID := uint64(1)
	addr, _, err := contractKeeper.Instantiate(ctx, codeID, funder, funder, initMsgBz, "demo contract", nil)
	require.NoError(t, err)

	return addr
}

func fundAccount(t *testing.T, ctx sdk.Context, tokenz *app.App, addr sdk.AccAddress, coins sdk.Coins) {
	t.Helper()
	// TODO:
	// err := simapp.FundAccount(
	// 	tokenz.BankKeeper,
	// 	ctx,
	// 	addr,
	// 	coins,
	// )

	// require.NoError(t, err)
	err := tokenz.BankKeeper.MintCoins(ctx, minttypes.ModuleName, coins)
	require.NoError(t, err)
	err = tokenz.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, coins)
	require.NoError(t, err)
}

func SetupCustomApp(t *testing.T, addr sdk.AccAddress) (*app.App, sdk.Context) {
	t.Helper()
	tokenz, ctx := CreateTestInput()
	wasmKeeper := tokenz.WasmKeeper

	storeReflectCode(t, ctx, tokenz, addr)

	cInfo := wasmKeeper.GetCodeInfo(ctx, 1)
	require.NotNil(t, cInfo)

	return tokenz, ctx
}
