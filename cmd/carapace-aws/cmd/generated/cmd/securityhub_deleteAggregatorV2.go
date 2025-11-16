package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_deleteAggregatorV2Cmd = &cobra.Command{
	Use:   "delete-aggregator-v2",
	Short: "Deletes the Aggregator V2.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_deleteAggregatorV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_deleteAggregatorV2Cmd).Standalone()

		securityhub_deleteAggregatorV2Cmd.Flags().String("aggregator-v2-arn", "", "The ARN of the Aggregator V2.")
		securityhub_deleteAggregatorV2Cmd.MarkFlagRequired("aggregator-v2-arn")
	})
	securityhubCmd.AddCommand(securityhub_deleteAggregatorV2Cmd)
}
