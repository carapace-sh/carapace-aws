package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisCmd = &cobra.Command{
	Use:   "kinesis",
	Short: "Amazon Kinesis Data Streams Service API Reference\n\nAmazon Kinesis Data Streams is a managed service that scales elastically for real-time processing of streaming big data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisCmd).Standalone()

	})
	rootCmd.AddCommand(kinesisCmd)
}
