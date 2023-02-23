package upgrade_test

import (
	errorsmod "cosmossdk.io/errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	stakeibckeeper "github.com/soohoio/stayking/x/stakeibc/keeper"
	stakeibcmigration "github.com/soohoio/stayking/x/stakeibc/migrations/upgrade_test"
	stakeibctypes "github.com/soohoio/stayking/x/stakeibc/types"
)

// Note: ensure these values are properly set before running upgrade
var (
	UpgradeName = "upgrade_test"
)

// Helper function to log the migrated modules consensus versions
func logModuleMigration(ctx sdk.Context, versionMap module.VersionMap, moduleName string) {
	currentVersion := versionMap[moduleName]
	ctx.Logger().Info(fmt.Sprintf("migrating module %s from version %d to version %d", moduleName, currentVersion-1, currentVersion))
}

// CreateUpgradeHandler creates an SDK upgrade handler for v5
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	cdc codec.Codec,
	stakeibcKeeper stakeibckeeper.Keeper,
	stakeibcStoreKey storetypes.StoreKey,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		currentVersions := mm.GetVersionMap()

		logModuleMigration(ctx, currentVersions, stakeibctypes.ModuleName)
		if err := stakeibcmigration.MigrateStore(ctx, stakeibcStoreKey, cdc); err != nil {
			return vm, errorsmod.Wrapf(err, "unable to migrate stakeibc store")
		}

		// `RunMigrations` (below) checks the old consensus version of each module (found in
		// the store) and compares it against the updated consensus version in the binary
		// If the old and new consensus versions are not the same, it attempts to call that
		// module's migration function that must be registered ahead of time
		//
		// Since the migrations above were executed directly (instead of being registered
		// and invoked through a Migrator), we need to set the module versions in the versionMap
		// to the new version, to prevent RunMigrations from attempting to re-run each migrations
		//vm[claimtypes.ModuleName] = currentVersions[claimtypes.ModuleName]
		//vm[icacallbacktypes.ModuleName] = currentVersions[icacallbacktypes.ModuleName]
		//vm[recordtypes.ModuleName] = currentVersions[recordtypes.ModuleName]
		vm[stakeibctypes.ModuleName] = currentVersions[stakeibctypes.ModuleName]

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
