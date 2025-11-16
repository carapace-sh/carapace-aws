package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listNodesSummaryCmd = &cobra.Command{
	Use:   "list-nodes-summary",
	Short: "Generates a summary of managed instance/node metadata based on the filters and aggregators you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listNodesSummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_listNodesSummaryCmd).Standalone()

		ssm_listNodesSummaryCmd.Flags().String("aggregators", "", "Specify one or more aggregators to return a count of managed nodes that match that expression.")
		ssm_listNodesSummaryCmd.Flags().String("filters", "", "One or more filters.")
		ssm_listNodesSummaryCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_listNodesSummaryCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_listNodesSummaryCmd.Flags().String("sync-name", "", "The name of the Amazon Web Services managed resource data sync to retrieve information about.")
		ssm_listNodesSummaryCmd.MarkFlagRequired("aggregators")
	})
	ssmCmd.AddCommand(ssm_listNodesSummaryCmd)
}
