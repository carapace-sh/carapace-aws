package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listSessionsCmd = &cobra.Command{
	Use:   "list-sessions",
	Short: "Lists sessions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listSessionsCmd).Standalone()

	deadline_listSessionsCmd.Flags().String("farm-id", "", "The farm ID for the list of sessions.")
	deadline_listSessionsCmd.Flags().String("job-id", "", "The job ID for the list of sessions.")
	deadline_listSessionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_listSessionsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_listSessionsCmd.Flags().String("queue-id", "", "The queue ID for the list of sessions")
	deadline_listSessionsCmd.MarkFlagRequired("farm-id")
	deadline_listSessionsCmd.MarkFlagRequired("job-id")
	deadline_listSessionsCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_listSessionsCmd)
}
