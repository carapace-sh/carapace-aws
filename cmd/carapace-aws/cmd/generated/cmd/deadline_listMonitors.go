package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listMonitorsCmd = &cobra.Command{
	Use:   "list-monitors",
	Short: "Gets a list of your monitors in Deadline Cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listMonitorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_listMonitorsCmd).Standalone()

		deadline_listMonitorsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		deadline_listMonitorsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	})
	deadlineCmd.AddCommand(deadline_listMonitorsCmd)
}
