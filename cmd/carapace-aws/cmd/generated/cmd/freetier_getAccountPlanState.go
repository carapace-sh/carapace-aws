package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var freetier_getAccountPlanStateCmd = &cobra.Command{
	Use:   "get-account-plan-state",
	Short: "This returns all of the information related to the state of the account plan related to Free Tier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(freetier_getAccountPlanStateCmd).Standalone()

	freetierCmd.AddCommand(freetier_getAccountPlanStateCmd)
}
