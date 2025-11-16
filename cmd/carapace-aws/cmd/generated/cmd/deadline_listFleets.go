package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listFleetsCmd = &cobra.Command{
	Use:   "list-fleets",
	Short: "Lists fleets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listFleetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listFleetsCmd).Standalone()

		deadline_listFleetsCmd.Flags().String("display-name", "", "The display names of a list of fleets.")
		deadline_listFleetsCmd.Flags().String("farm-id", "", "The farm ID of the fleets.")
		deadline_listFleetsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listFleetsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
		deadline_listFleetsCmd.Flags().String("principal-id", "", "The principal ID of the members to include in the fleet.")
		deadline_listFleetsCmd.Flags().String("status", "", "The status of the fleet.")
		deadline_listFleetsCmd.MarkFlagRequired("farm-id")
	})
	deadlineCmd.AddCommand(deadline_listFleetsCmd)
}
