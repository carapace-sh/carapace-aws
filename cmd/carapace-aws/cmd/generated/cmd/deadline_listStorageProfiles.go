package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listStorageProfilesCmd = &cobra.Command{
	Use:   "list-storage-profiles",
	Short: "Lists storage profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listStorageProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listStorageProfilesCmd).Standalone()

		deadline_listStorageProfilesCmd.Flags().String("farm-id", "", "The farm ID of the storage profile.")
		deadline_listStorageProfilesCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listStorageProfilesCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listStorageProfilesCmd.MarkFlagRequired("farm-id")
	})
	deadlineCmd.AddCommand(deadline_listStorageProfilesCmd)
}
