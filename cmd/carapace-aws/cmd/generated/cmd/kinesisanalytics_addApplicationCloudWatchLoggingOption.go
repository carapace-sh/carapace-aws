package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_addApplicationCloudWatchLoggingOptionCmd = &cobra.Command{
	Use:   "add-application-cloud-watch-logging-option",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_addApplicationCloudWatchLoggingOptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalytics_addApplicationCloudWatchLoggingOptionCmd).Standalone()

		kinesisanalytics_addApplicationCloudWatchLoggingOptionCmd.Flags().String("application-name", "", "The Kinesis Analytics application name.")
		kinesisanalytics_addApplicationCloudWatchLoggingOptionCmd.Flags().String("cloud-watch-logging-option", "", "Provides the CloudWatch log stream Amazon Resource Name (ARN) and the IAM role ARN.")
		kinesisanalytics_addApplicationCloudWatchLoggingOptionCmd.Flags().String("current-application-version-id", "", "The version ID of the Kinesis Analytics application.")
		kinesisanalytics_addApplicationCloudWatchLoggingOptionCmd.MarkFlagRequired("application-name")
		kinesisanalytics_addApplicationCloudWatchLoggingOptionCmd.MarkFlagRequired("cloud-watch-logging-option")
		kinesisanalytics_addApplicationCloudWatchLoggingOptionCmd.MarkFlagRequired("current-application-version-id")
	})
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_addApplicationCloudWatchLoggingOptionCmd)
}
