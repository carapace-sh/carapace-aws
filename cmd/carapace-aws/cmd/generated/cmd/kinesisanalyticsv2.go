package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2Cmd = &cobra.Command{
	Use:   "kinesisanalyticsv2",
	Short: "Amazon Managed Service for Apache Flink was previously known as Amazon Kinesis Data Analytics for Apache Flink.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2Cmd).Standalone()

	})
	rootCmd.AddCommand(kinesisanalyticsv2Cmd)
}
