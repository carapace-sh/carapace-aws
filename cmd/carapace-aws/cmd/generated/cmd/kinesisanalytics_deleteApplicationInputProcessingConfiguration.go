package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_deleteApplicationInputProcessingConfigurationCmd = &cobra.Command{
	Use:   "delete-application-input-processing-configuration",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_deleteApplicationInputProcessingConfigurationCmd).Standalone()

	kinesisanalytics_deleteApplicationInputProcessingConfigurationCmd.Flags().String("application-name", "", "The Kinesis Analytics application name.")
	kinesisanalytics_deleteApplicationInputProcessingConfigurationCmd.Flags().String("current-application-version-id", "", "The version ID of the Kinesis Analytics application.")
	kinesisanalytics_deleteApplicationInputProcessingConfigurationCmd.Flags().String("input-id", "", "The ID of the input configuration from which to delete the input processing configuration.")
	kinesisanalytics_deleteApplicationInputProcessingConfigurationCmd.MarkFlagRequired("application-name")
	kinesisanalytics_deleteApplicationInputProcessingConfigurationCmd.MarkFlagRequired("current-application-version-id")
	kinesisanalytics_deleteApplicationInputProcessingConfigurationCmd.MarkFlagRequired("input-id")
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_deleteApplicationInputProcessingConfigurationCmd)
}
