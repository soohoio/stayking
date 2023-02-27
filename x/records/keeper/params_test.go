package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/soohoio/stayking/v2/testutil/keeper"
	"github.com/soohoio/stayking/v2/x/records/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RecordsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
