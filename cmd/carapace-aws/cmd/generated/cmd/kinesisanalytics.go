package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsCmd = &cobra.Command{
	Use:   "kinesisanalytics",
	Short: "Amazon Kinesis Analytics\n\n**Overview**\n\nThis documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsCmd).Standalone()

	})
	rootCmd.AddCommand(kinesisanalyticsCmd)
}
