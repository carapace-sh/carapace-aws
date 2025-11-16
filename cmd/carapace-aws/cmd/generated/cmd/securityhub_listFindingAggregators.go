package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listFindingAggregatorsCmd = &cobra.Command{
	Use:   "list-finding-aggregators",
	Short: "If cross-Region aggregation is enabled, then `ListFindingAggregators` returns the Amazon Resource Name (ARN) of the finding aggregator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listFindingAggregatorsCmd).Standalone()

	securityhub_listFindingAggregatorsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	securityhub_listFindingAggregatorsCmd.Flags().String("next-token", "", "The token returned with the previous set of results.")
	securityhubCmd.AddCommand(securityhub_listFindingAggregatorsCmd)
}
