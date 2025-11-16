package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_listStreamConsumersCmd = &cobra.Command{
	Use:   "list-stream-consumers",
	Short: "Lists the consumers registered to receive data from a stream using enhanced fan-out, and provides information about each consumer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_listStreamConsumersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_listStreamConsumersCmd).Standalone()

		kinesis_listStreamConsumersCmd.Flags().String("max-results", "", "The maximum number of consumers that you want a single call of `ListStreamConsumers` to return.")
		kinesis_listStreamConsumersCmd.Flags().String("next-token", "", "When the number of consumers that are registered with the data stream is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of consumers that are registered with the data stream, the response includes a pagination token named `NextToken`.")
		kinesis_listStreamConsumersCmd.Flags().String("stream-arn", "", "The ARN of the Kinesis data stream for which you want to list the registered consumers.")
		kinesis_listStreamConsumersCmd.Flags().String("stream-creation-timestamp", "", "Specify this input parameter to distinguish data streams that have the same name.")
		kinesis_listStreamConsumersCmd.MarkFlagRequired("stream-arn")
	})
	kinesisCmd.AddCommand(kinesis_listStreamConsumersCmd)
}
