package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_batchGetAggregateResourceConfigCmd = &cobra.Command{
	Use:   "batch-get-aggregate-resource-config",
	Short: "Returns the current configuration items for resources that are present in your Config aggregator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_batchGetAggregateResourceConfigCmd).Standalone()

	config_batchGetAggregateResourceConfigCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
	config_batchGetAggregateResourceConfigCmd.Flags().String("resource-identifiers", "", "A list of aggregate ResourceIdentifiers objects.")
	config_batchGetAggregateResourceConfigCmd.MarkFlagRequired("configuration-aggregator-name")
	config_batchGetAggregateResourceConfigCmd.MarkFlagRequired("resource-identifiers")
	configCmd.AddCommand(config_batchGetAggregateResourceConfigCmd)
}
