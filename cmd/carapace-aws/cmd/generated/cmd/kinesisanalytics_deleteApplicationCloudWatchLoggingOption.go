package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_deleteApplicationCloudWatchLoggingOptionCmd = &cobra.Command{
	Use:   "delete-application-cloud-watch-logging-option",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_deleteApplicationCloudWatchLoggingOptionCmd).Standalone()

	kinesisanalytics_deleteApplicationCloudWatchLoggingOptionCmd.Flags().String("application-name", "", "The Kinesis Analytics application name.")
	kinesisanalytics_deleteApplicationCloudWatchLoggingOptionCmd.Flags().String("cloud-watch-logging-option-id", "", "The `CloudWatchLoggingOptionId` of the CloudWatch logging option to delete.")
	kinesisanalytics_deleteApplicationCloudWatchLoggingOptionCmd.Flags().String("current-application-version-id", "", "The version ID of the Kinesis Analytics application.")
	kinesisanalytics_deleteApplicationCloudWatchLoggingOptionCmd.MarkFlagRequired("application-name")
	kinesisanalytics_deleteApplicationCloudWatchLoggingOptionCmd.MarkFlagRequired("cloud-watch-logging-option-id")
	kinesisanalytics_deleteApplicationCloudWatchLoggingOptionCmd.MarkFlagRequired("current-application-version-id")
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_deleteApplicationCloudWatchLoggingOptionCmd)
}
