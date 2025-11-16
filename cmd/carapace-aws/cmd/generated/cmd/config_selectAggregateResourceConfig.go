package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_selectAggregateResourceConfigCmd = &cobra.Command{
	Use:   "select-aggregate-resource-config",
	Short: "Accepts a structured query language (SQL) SELECT command and an aggregator to query configuration state of Amazon Web Services resources across multiple accounts and regions, performs the corresponding search, and returns resource configurations matching the properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_selectAggregateResourceConfigCmd).Standalone()

	config_selectAggregateResourceConfigCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
	config_selectAggregateResourceConfigCmd.Flags().String("expression", "", "The SQL query SELECT command.")
	config_selectAggregateResourceConfigCmd.Flags().String("limit", "", "The maximum number of query results returned on each page.")
	config_selectAggregateResourceConfigCmd.Flags().String("max-results", "", "The maximum number of query results returned on each page.")
	config_selectAggregateResourceConfigCmd.Flags().String("next-token", "", "The nextToken string returned in a previous request that you use to request the next page of results in a paginated response.")
	config_selectAggregateResourceConfigCmd.MarkFlagRequired("configuration-aggregator-name")
	config_selectAggregateResourceConfigCmd.MarkFlagRequired("expression")
	configCmd.AddCommand(config_selectAggregateResourceConfigCmd)
}
