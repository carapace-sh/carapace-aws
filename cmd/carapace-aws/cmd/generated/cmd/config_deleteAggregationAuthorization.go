package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_deleteAggregationAuthorizationCmd = &cobra.Command{
	Use:   "delete-aggregation-authorization",
	Short: "Deletes the authorization granted to the specified configuration aggregator account in a specified region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteAggregationAuthorizationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_deleteAggregationAuthorizationCmd).Standalone()

		config_deleteAggregationAuthorizationCmd.Flags().String("authorized-account-id", "", "The 12-digit account ID of the account authorized to aggregate data.")
		config_deleteAggregationAuthorizationCmd.Flags().String("authorized-aws-region", "", "The region authorized to collect aggregated data.")
		config_deleteAggregationAuthorizationCmd.MarkFlagRequired("authorized-account-id")
		config_deleteAggregationAuthorizationCmd.MarkFlagRequired("authorized-aws-region")
	})
	configCmd.AddCommand(config_deleteAggregationAuthorizationCmd)
}
