package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var savingsplans_returnSavingsPlanCmd = &cobra.Command{
	Use:   "return-savings-plan",
	Short: "Returns the specified Savings Plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(savingsplans_returnSavingsPlanCmd).Standalone()

	savingsplans_returnSavingsPlanCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	savingsplans_returnSavingsPlanCmd.Flags().String("savings-plan-id", "", "The ID of the Savings Plan.")
	savingsplans_returnSavingsPlanCmd.MarkFlagRequired("savings-plan-id")
	savingsplansCmd.AddCommand(savingsplans_returnSavingsPlanCmd)
}
