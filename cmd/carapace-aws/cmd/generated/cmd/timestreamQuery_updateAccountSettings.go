package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_updateAccountSettingsCmd = &cobra.Command{
	Use:   "update-account-settings",
	Short: "Transitions your account to use TCUs for query pricing and modifies the maximum query compute units that you've configured.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_updateAccountSettingsCmd).Standalone()

	timestreamQuery_updateAccountSettingsCmd.Flags().String("max-query-tcu", "", "The maximum number of compute units the service will use at any point in time to serve your queries.")
	timestreamQuery_updateAccountSettingsCmd.Flags().String("query-compute", "", "Modifies the query compute settings configured in your account, including the query pricing model and provisioned Timestream Compute Units (TCUs) in your account.")
	timestreamQuery_updateAccountSettingsCmd.Flags().String("query-pricing-model", "", "The pricing model for queries in an account.")
	timestreamQueryCmd.AddCommand(timestreamQuery_updateAccountSettingsCmd)
}
