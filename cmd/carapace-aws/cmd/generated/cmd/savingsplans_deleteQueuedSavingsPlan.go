package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var savingsplans_deleteQueuedSavingsPlanCmd = &cobra.Command{
	Use:   "delete-queued-savings-plan",
	Short: "Deletes the queued purchase for the specified Savings Plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(savingsplans_deleteQueuedSavingsPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(savingsplans_deleteQueuedSavingsPlanCmd).Standalone()

		savingsplans_deleteQueuedSavingsPlanCmd.Flags().String("savings-plan-id", "", "The ID of the Savings Plan.")
		savingsplans_deleteQueuedSavingsPlanCmd.MarkFlagRequired("savings-plan-id")
	})
	savingsplansCmd.AddCommand(savingsplans_deleteQueuedSavingsPlanCmd)
}
