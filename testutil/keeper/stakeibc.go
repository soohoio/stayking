package keeper

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	strideapp "github.com/soohoio/stayking/v2/app"
	"github.com/soohoio/stayking/v2/x/stakeibc/keeper"
)

func StakeibcKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	app := strideapp.InitStayKingTestApp(true)
	stakeibcKeeper := app.StakeibcKeeper
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1, ChainID: "stayking-1", Time: time.Now().UTC()})

	return &stakeibcKeeper, ctx
}
