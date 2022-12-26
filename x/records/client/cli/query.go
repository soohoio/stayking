package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/soohoio/stayking/x/records/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group records queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdListUserRedemptionRecord())
	cmd.AddCommand(CmdShowUserRedemptionRecord())
	cmd.AddCommand(CmdListEpochUnbondingRecord())
	cmd.AddCommand(CmdShowEpochUnbondingRecord())
	cmd.AddCommand(CmdListDepositRecord())
	cmd.AddCommand(CmdShowDepositRecord())
	// this line is used by starport scaffolding # 1

	return cmd
}
