package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeConfigurationAggregatorSourcesStatusCmd = &cobra.Command{
	Use:   "describe-configuration-aggregator-sources-status",
	Short: "Returns status information for sources within an aggregator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeConfigurationAggregatorSourcesStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeConfigurationAggregatorSourcesStatusCmd).Standalone()

		config_describeConfigurationAggregatorSourcesStatusCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
		config_describeConfigurationAggregatorSourcesStatusCmd.Flags().String("limit", "", "The maximum number of AggregatorSourceStatus returned on each page.")
		config_describeConfigurationAggregatorSourcesStatusCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_describeConfigurationAggregatorSourcesStatusCmd.Flags().String("update-status", "", "Filters the status type.")
		config_describeConfigurationAggregatorSourcesStatusCmd.MarkFlagRequired("configuration-aggregator-name")
	})
	configCmd.AddCommand(config_describeConfigurationAggregatorSourcesStatusCmd)
}
