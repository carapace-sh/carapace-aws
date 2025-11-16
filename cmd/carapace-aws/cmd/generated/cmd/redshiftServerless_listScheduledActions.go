package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftServerless_listScheduledActionsCmd = &cobra.Command{
	Use:   "list-scheduled-actions",
	Short: "Returns a list of scheduled actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftServerless_listScheduledActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftServerless_listScheduledActionsCmd).Standalone()

		redshiftServerless_listScheduledActionsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		redshiftServerless_listScheduledActionsCmd.Flags().String("namespace-name", "", "The name of namespace associated with the scheduled action to retrieve.")
		redshiftServerless_listScheduledActionsCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	})
	redshiftServerlessCmd.AddCommand(redshiftServerless_listScheduledActionsCmd)
}
