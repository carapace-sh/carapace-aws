package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_getAggregatorV2Cmd = &cobra.Command{
	Use:   "get-aggregator-v2",
	Short: "Returns the configuration of the specified Aggregator V2.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_getAggregatorV2Cmd).Standalone()

	securityhub_getAggregatorV2Cmd.Flags().String("aggregator-v2-arn", "", "The ARN of the Aggregator V2.")
	securityhub_getAggregatorV2Cmd.MarkFlagRequired("aggregator-v2-arn")
	securityhubCmd.AddCommand(securityhub_getAggregatorV2Cmd)
}
