package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteFarmCmd = &cobra.Command{
	Use:   "delete-farm",
	Short: "Deletes a farm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteFarmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_deleteFarmCmd).Standalone()

		deadline_deleteFarmCmd.Flags().String("farm-id", "", "The farm ID of the farm to delete.")
		deadline_deleteFarmCmd.MarkFlagRequired("farm-id")
	})
	deadlineCmd.AddCommand(deadline_deleteFarmCmd)
}
