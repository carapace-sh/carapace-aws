package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putConfigurationAggregatorCmd = &cobra.Command{
	Use:   "put-configuration-aggregator",
	Short: "Creates and updates the configuration aggregator with the selected source accounts and regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putConfigurationAggregatorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_putConfigurationAggregatorCmd).Standalone()

		config_putConfigurationAggregatorCmd.Flags().String("account-aggregation-sources", "", "A list of AccountAggregationSource object.")
		config_putConfigurationAggregatorCmd.Flags().String("aggregator-filters", "", "An object to filter configuration recorders in an aggregator.")
		config_putConfigurationAggregatorCmd.Flags().String("configuration-aggregator-name", "", "The name of the configuration aggregator.")
		config_putConfigurationAggregatorCmd.Flags().String("organization-aggregation-source", "", "An OrganizationAggregationSource object.")
		config_putConfigurationAggregatorCmd.Flags().String("tags", "", "An array of tag object.")
		config_putConfigurationAggregatorCmd.MarkFlagRequired("configuration-aggregator-name")
	})
	configCmd.AddCommand(config_putConfigurationAggregatorCmd)
}
