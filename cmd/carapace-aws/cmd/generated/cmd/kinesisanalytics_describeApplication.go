package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_describeApplicationCmd = &cobra.Command{
	Use:   "describe-application",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_describeApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalytics_describeApplicationCmd).Standalone()

		kinesisanalytics_describeApplicationCmd.Flags().String("application-name", "", "Name of the application.")
		kinesisanalytics_describeApplicationCmd.MarkFlagRequired("application-name")
	})
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_describeApplicationCmd)
}
