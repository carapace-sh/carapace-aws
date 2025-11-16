package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_deleteApplicationOutputCmd = &cobra.Command{
	Use:   "delete-application-output",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_deleteApplicationOutputCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalytics_deleteApplicationOutputCmd).Standalone()

		kinesisanalytics_deleteApplicationOutputCmd.Flags().String("application-name", "", "Amazon Kinesis Analytics application name.")
		kinesisanalytics_deleteApplicationOutputCmd.Flags().String("current-application-version-id", "", "Amazon Kinesis Analytics application version.")
		kinesisanalytics_deleteApplicationOutputCmd.Flags().String("output-id", "", "The ID of the configuration to delete.")
		kinesisanalytics_deleteApplicationOutputCmd.MarkFlagRequired("application-name")
		kinesisanalytics_deleteApplicationOutputCmd.MarkFlagRequired("current-application-version-id")
		kinesisanalytics_deleteApplicationOutputCmd.MarkFlagRequired("output-id")
	})
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_deleteApplicationOutputCmd)
}
