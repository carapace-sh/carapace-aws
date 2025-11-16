package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeConfigurationAggregatorsCmd = &cobra.Command{
	Use:   "describe-configuration-aggregators",
	Short: "Returns the details of one or more configuration aggregators.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeConfigurationAggregatorsCmd).Standalone()

	config_describeConfigurationAggregatorsCmd.Flags().String("configuration-aggregator-names", "", "The name of the configuration aggregators.")
	config_describeConfigurationAggregatorsCmd.Flags().String("limit", "", "The maximum number of configuration aggregators returned on each page.")
	config_describeConfigurationAggregatorsCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
	configCmd.AddCommand(config_describeConfigurationAggregatorsCmd)
}
