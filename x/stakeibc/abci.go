package stakeibc

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"

	"github.com/soohoio/stayking/v2/x/stakeibc/keeper"
	"github.com/soohoio/stayking/v2/x/stakeibc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlocker of stakeibc module
func BeginBlocker(ctx sdk.Context, k keeper.Keeper, bk types.BankKeeper, ak types.AccountKeeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	// Iterate over all host zones and verify redemption rate
	//for _, hz := range k.GetAllHostZone(ctx) {
	//	rrSafe, err := k.IsRedemptionRateWithinSafetyBounds(ctx, hz)
	//	if !rrSafe {
	//		panic(fmt.Sprintf("[INVARIANT BROKEN!!!] %s's RR is %s. ERR: %v", hz.GetChainId(), hz.RedemptionRate.String(), err.Error()))
	//	}
	//}
}
