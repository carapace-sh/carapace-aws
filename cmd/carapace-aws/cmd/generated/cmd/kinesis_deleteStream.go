package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_deleteStreamCmd = &cobra.Command{
	Use:   "delete-stream",
	Short: "Deletes a Kinesis data stream and all its shards and data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_deleteStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_deleteStreamCmd).Standalone()

		kinesis_deleteStreamCmd.Flags().String("enforce-consumer-deletion", "", "If this parameter is unset (`null`) or if you set it to `false`, and the stream has registered consumers, the call to `DeleteStream` fails with a `ResourceInUseException`.")
		kinesis_deleteStreamCmd.Flags().String("stream-arn", "", "The ARN of the stream.")
		kinesis_deleteStreamCmd.Flags().String("stream-name", "", "The name of the stream to delete.")
	})
	kinesisCmd.AddCommand(kinesis_deleteStreamCmd)
}
