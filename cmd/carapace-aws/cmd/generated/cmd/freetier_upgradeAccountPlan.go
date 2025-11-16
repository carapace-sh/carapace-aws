package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var freetier_upgradeAccountPlanCmd = &cobra.Command{
	Use:   "upgrade-account-plan",
	Short: "The account plan type for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(freetier_upgradeAccountPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(freetier_upgradeAccountPlanCmd).Standalone()

		freetier_upgradeAccountPlanCmd.Flags().String("account-plan-type", "", "The target account plan type.")
		freetier_upgradeAccountPlanCmd.MarkFlagRequired("account-plan-type")
	})
	freetierCmd.AddCommand(freetier_upgradeAccountPlanCmd)
}
