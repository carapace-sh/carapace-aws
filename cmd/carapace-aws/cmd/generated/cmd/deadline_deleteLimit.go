package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteLimitCmd = &cobra.Command{
	Use:   "delete-limit",
	Short: "Removes a limit from the specified farm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteLimitCmd).Standalone()

	deadline_deleteLimitCmd.Flags().String("farm-id", "", "The unique identifier of the farm that contains the limit to delete.")
	deadline_deleteLimitCmd.Flags().String("limit-id", "", "The unique identifier of the limit to delete.")
	deadline_deleteLimitCmd.MarkFlagRequired("farm-id")
	deadline_deleteLimitCmd.MarkFlagRequired("limit-id")
	deadlineCmd.AddCommand(deadline_deleteLimitCmd)
}
