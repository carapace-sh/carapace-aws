package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listWorkersCmd = &cobra.Command{
	Use:   "list-workers",
	Short: "Lists workers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listWorkersCmd).Standalone()

	deadline_listWorkersCmd.Flags().String("farm-id", "", "The farm ID connected to the workers.")
	deadline_listWorkersCmd.Flags().String("fleet-id", "", "The fleet ID of the workers.")
	deadline_listWorkersCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_listWorkersCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_listWorkersCmd.MarkFlagRequired("farm-id")
	deadline_listWorkersCmd.MarkFlagRequired("fleet-id")
	deadlineCmd.AddCommand(deadline_listWorkersCmd)
}
