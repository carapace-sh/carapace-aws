package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_updateAggregatorV2Cmd = &cobra.Command{
	Use:   "update-aggregator-v2",
	Short: "Udpates the configuration for the Aggregator V2.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_updateAggregatorV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_updateAggregatorV2Cmd).Standalone()

		securityhub_updateAggregatorV2Cmd.Flags().String("aggregator-v2-arn", "", "The ARN of the Aggregator V2.")
		securityhub_updateAggregatorV2Cmd.Flags().String("linked-regions", "", "A list of Amazon Web Services Regions linked to the aggegation Region.")
		securityhub_updateAggregatorV2Cmd.Flags().String("region-linking-mode", "", "Determines how Amazon Web Services Regions should be linked to the Aggregator V2.")
		securityhub_updateAggregatorV2Cmd.MarkFlagRequired("aggregator-v2-arn")
		securityhub_updateAggregatorV2Cmd.MarkFlagRequired("region-linking-mode")
	})
	securityhubCmd.AddCommand(securityhub_updateAggregatorV2Cmd)
}
