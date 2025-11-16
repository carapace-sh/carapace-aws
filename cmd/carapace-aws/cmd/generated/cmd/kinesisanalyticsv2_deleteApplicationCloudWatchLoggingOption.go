package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_deleteApplicationCloudWatchLoggingOptionCmd = &cobra.Command{
	Use:   "delete-application-cloud-watch-logging-option",
	Short: "Deletes an Amazon CloudWatch log stream from an SQL-based Kinesis Data Analytics application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_deleteApplicationCloudWatchLoggingOptionCmd).Standalone()

	kinesisanalyticsv2_deleteApplicationCloudWatchLoggingOptionCmd.Flags().String("application-name", "", "The application name.")
	kinesisanalyticsv2_deleteApplicationCloudWatchLoggingOptionCmd.Flags().String("cloud-watch-logging-option-id", "", "The `CloudWatchLoggingOptionId` of the Amazon CloudWatch logging option to delete.")
	kinesisanalyticsv2_deleteApplicationCloudWatchLoggingOptionCmd.Flags().String("conditional-token", "", "A value you use to implement strong concurrency for application updates.")
	kinesisanalyticsv2_deleteApplicationCloudWatchLoggingOptionCmd.Flags().String("current-application-version-id", "", "The version ID of the application.")
	kinesisanalyticsv2_deleteApplicationCloudWatchLoggingOptionCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2_deleteApplicationCloudWatchLoggingOptionCmd.MarkFlagRequired("cloud-watch-logging-option-id")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_deleteApplicationCloudWatchLoggingOptionCmd)
}
