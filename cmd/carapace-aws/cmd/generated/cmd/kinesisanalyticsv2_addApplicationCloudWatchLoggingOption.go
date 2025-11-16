package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_addApplicationCloudWatchLoggingOptionCmd = &cobra.Command{
	Use:   "add-application-cloud-watch-logging-option",
	Short: "Adds an Amazon CloudWatch log stream to monitor application configuration errors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_addApplicationCloudWatchLoggingOptionCmd).Standalone()

	kinesisanalyticsv2_addApplicationCloudWatchLoggingOptionCmd.Flags().String("application-name", "", "The Kinesis Data Analytics application name.")
	kinesisanalyticsv2_addApplicationCloudWatchLoggingOptionCmd.Flags().String("cloud-watch-logging-option", "", "Provides the Amazon CloudWatch log stream Amazon Resource Name (ARN).")
	kinesisanalyticsv2_addApplicationCloudWatchLoggingOptionCmd.Flags().String("conditional-token", "", "A value you use to implement strong concurrency for application updates.")
	kinesisanalyticsv2_addApplicationCloudWatchLoggingOptionCmd.Flags().String("current-application-version-id", "", "The version ID of the SQL-based Kinesis Data Analytics application.")
	kinesisanalyticsv2_addApplicationCloudWatchLoggingOptionCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2_addApplicationCloudWatchLoggingOptionCmd.MarkFlagRequired("cloud-watch-logging-option")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_addApplicationCloudWatchLoggingOptionCmd)
}
