package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listFindingAggregationsCmd = &cobra.Command{
	Use:   "list-finding-aggregations",
	Short: "Lists aggregated finding data for your environment based on specific criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listFindingAggregationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_listFindingAggregationsCmd).Standalone()

		inspector2_listFindingAggregationsCmd.Flags().String("account-ids", "", "The Amazon Web Services account IDs to retrieve finding aggregation data for.")
		inspector2_listFindingAggregationsCmd.Flags().String("aggregation-request", "", "Details of the aggregation request that is used to filter your aggregation results.")
		inspector2_listFindingAggregationsCmd.Flags().String("aggregation-type", "", "The type of the aggregation request.")
		inspector2_listFindingAggregationsCmd.Flags().String("max-results", "", "The maximum number of results the response can return.")
		inspector2_listFindingAggregationsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
		inspector2_listFindingAggregationsCmd.MarkFlagRequired("aggregation-type")
	})
	inspector2Cmd.AddCommand(inspector2_listFindingAggregationsCmd)
}
