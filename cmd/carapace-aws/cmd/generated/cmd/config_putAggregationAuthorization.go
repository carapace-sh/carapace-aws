package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_putAggregationAuthorizationCmd = &cobra.Command{
	Use:   "put-aggregation-authorization",
	Short: "Authorizes the aggregator account and region to collect data from the source account and region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_putAggregationAuthorizationCmd).Standalone()

	config_putAggregationAuthorizationCmd.Flags().String("authorized-account-id", "", "The 12-digit account ID of the account authorized to aggregate data.")
	config_putAggregationAuthorizationCmd.Flags().String("authorized-aws-region", "", "The region authorized to collect aggregated data.")
	config_putAggregationAuthorizationCmd.Flags().String("tags", "", "An array of tag object.")
	config_putAggregationAuthorizationCmd.MarkFlagRequired("authorized-account-id")
	config_putAggregationAuthorizationCmd.MarkFlagRequired("authorized-aws-region")
	configCmd.AddCommand(config_putAggregationAuthorizationCmd)
}
