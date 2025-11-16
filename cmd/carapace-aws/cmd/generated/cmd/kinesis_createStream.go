package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_createStreamCmd = &cobra.Command{
	Use:   "create-stream",
	Short: "Creates a Kinesis data stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_createStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesis_createStreamCmd).Standalone()

		kinesis_createStreamCmd.Flags().String("max-record-size-in-ki-b", "", "The maximum record size of a single record in kibibyte (KiB) that you can write to, and read from a stream.")
		kinesis_createStreamCmd.Flags().String("shard-count", "", "The number of shards that the stream will use.")
		kinesis_createStreamCmd.Flags().String("stream-mode-details", "", "Indicates the capacity mode of the data stream.")
		kinesis_createStreamCmd.Flags().String("stream-name", "", "A name to identify the stream.")
		kinesis_createStreamCmd.Flags().String("tags", "", "A set of up to 50 key-value pairs to use to create the tags.")
		kinesis_createStreamCmd.Flags().String("warm-throughput-mi-bps", "", "The target warm throughput in MB/s that the stream should be scaled to handle.")
		kinesis_createStreamCmd.MarkFlagRequired("stream-name")
	})
	kinesisCmd.AddCommand(kinesis_createStreamCmd)
}
