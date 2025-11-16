package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_createAggregatorV2Cmd = &cobra.Command{
	Use:   "create-aggregator-v2",
	Short: "Enables aggregation across Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_createAggregatorV2Cmd).Standalone()

	securityhub_createAggregatorV2Cmd.Flags().String("client-token", "", "A unique identifier used to ensure idempotency.")
	securityhub_createAggregatorV2Cmd.Flags().String("linked-regions", "", "The list of Regions that are linked to the aggregation Region.")
	securityhub_createAggregatorV2Cmd.Flags().String("region-linking-mode", "", "Determines how Regions are linked to an Aggregator V2.")
	securityhub_createAggregatorV2Cmd.Flags().String("tags", "", "A list of key-value pairs to be applied to the AggregatorV2.")
	securityhub_createAggregatorV2Cmd.MarkFlagRequired("region-linking-mode")
	securityhubCmd.AddCommand(securityhub_createAggregatorV2Cmd)
}
