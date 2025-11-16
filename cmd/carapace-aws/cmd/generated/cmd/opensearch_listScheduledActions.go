package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_listScheduledActionsCmd = &cobra.Command{
	Use:   "list-scheduled-actions",
	Short: "Retrieves a list of configuration changes that are scheduled for a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_listScheduledActionsCmd).Standalone()

	opensearch_listScheduledActionsCmd.Flags().String("domain-name", "", "The name of the domain.")
	opensearch_listScheduledActionsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	opensearch_listScheduledActionsCmd.Flags().String("next-token", "", "If your initial `ListScheduledActions` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `ListScheduledActions` operations, which returns results in the next page.")
	opensearch_listScheduledActionsCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_listScheduledActionsCmd)
}
