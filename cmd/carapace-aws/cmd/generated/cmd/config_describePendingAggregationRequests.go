package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describePendingAggregationRequestsCmd = &cobra.Command{
	Use:   "describe-pending-aggregation-requests",
	Short: "Returns a list of all pending aggregation requests.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describePendingAggregationRequestsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describePendingAggregationRequestsCmd).Standalone()

		config_describePendingAggregationRequestsCmd.Flags().String("limit", "", "The maximum number of evaluation results returned on each page.")
		config_describePendingAggregationRequestsCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	})
	configCmd.AddCommand(config_describePendingAggregationRequestsCmd)
}
