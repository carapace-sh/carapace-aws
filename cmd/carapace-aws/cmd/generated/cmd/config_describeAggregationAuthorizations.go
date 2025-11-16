package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeAggregationAuthorizationsCmd = &cobra.Command{
	Use:   "describe-aggregation-authorizations",
	Short: "Returns a list of authorizations granted to various aggregator accounts and regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeAggregationAuthorizationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeAggregationAuthorizationsCmd).Standalone()

		config_describeAggregationAuthorizationsCmd.Flags().String("limit", "", "The maximum number of AggregationAuthorizations returned on each page.")
		config_describeAggregationAuthorizationsCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	})
	configCmd.AddCommand(config_describeAggregationAuthorizationsCmd)
}
