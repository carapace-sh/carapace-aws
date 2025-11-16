package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateLimitCmd = &cobra.Command{
	Use:   "update-limit",
	Short: "Updates the properties of the specified limit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateLimitCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_updateLimitCmd).Standalone()

		deadline_updateLimitCmd.Flags().String("description", "", "The new description of the limit.")
		deadline_updateLimitCmd.Flags().String("display-name", "", "The new display name of the limit.")
		deadline_updateLimitCmd.Flags().String("farm-id", "", "The unique identifier of the farm that contains the limit.")
		deadline_updateLimitCmd.Flags().String("limit-id", "", "The unique identifier of the limit to update.")
		deadline_updateLimitCmd.Flags().String("max-count", "", "The maximum number of resources constrained by this limit.")
		deadline_updateLimitCmd.MarkFlagRequired("farm-id")
		deadline_updateLimitCmd.MarkFlagRequired("limit-id")
	})
	deadlineCmd.AddCommand(deadline_updateLimitCmd)
}
