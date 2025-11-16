package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_publishBatchCmd = &cobra.Command{
	Use:   "publish-batch",
	Short: "Publishes up to 10 messages to the specified topic in a single batch.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_publishBatchCmd).Standalone()

	sns_publishBatchCmd.Flags().String("publish-batch-request-entries", "", "A list of `PublishBatch` request entries to be sent to the SNS topic.")
	sns_publishBatchCmd.Flags().String("topic-arn", "", "The Amazon resource name (ARN) of the topic you want to batch publish to.")
	sns_publishBatchCmd.MarkFlagRequired("publish-batch-request-entries")
	sns_publishBatchCmd.MarkFlagRequired("topic-arn")
	snsCmd.AddCommand(sns_publishBatchCmd)
}
