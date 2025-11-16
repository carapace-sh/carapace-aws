package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_getAggregateResourceConfigCmd = &cobra.Command{
	Use:   "get-aggregate-resource-config",
	Short: "Returns configuration item that is aggregated for your specific resource in a specific source account and region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getAggregateResourceConfigCmd).Standalone()

	config_getAggregateResourceConfigCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
	config_getAggregateResourceConfigCmd.Flags().String("resource-identifier", "", "An object that identifies aggregate resource.")
	config_getAggregateResourceConfigCmd.MarkFlagRequired("configuration-aggregator-name")
	config_getAggregateResourceConfigCmd.MarkFlagRequired("resource-identifier")
	configCmd.AddCommand(config_getAggregateResourceConfigCmd)
}
