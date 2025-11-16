package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsCmd = &cobra.Command{
	Use:   "kinesisanalytics",
	Short: "Amazon Kinesis Analytics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsCmd).Standalone()

	})
	rootCmd.AddCommand(kinesisanalyticsCmd)
}
