package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteConfigurationAggregatorCmd = &cobra.Command{
	Use:   "delete-configuration-aggregator",
	Short: "Deletes the specified configuration aggregator and the aggregated data associated with the aggregator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteConfigurationAggregatorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_deleteConfigurationAggregatorCmd).Standalone()

		config_deleteConfigurationAggregatorCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
		config_deleteConfigurationAggregatorCmd.MarkFlagRequired("configuration-aggregator-name")
	})
	configCmd.AddCommand(config_deleteConfigurationAggregatorCmd)
}
