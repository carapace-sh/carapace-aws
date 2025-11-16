package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalytics_addApplicationInputProcessingConfigurationCmd = &cobra.Command{
	Use:   "add-application-input-processing-configuration",
	Short: "This documentation is for version 1 of the Amazon Kinesis Data Analytics API, which only supports SQL applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalytics_addApplicationInputProcessingConfigurationCmd).Standalone()

	kinesisanalytics_addApplicationInputProcessingConfigurationCmd.Flags().String("application-name", "", "Name of the application to which you want to add the input processing configuration.")
	kinesisanalytics_addApplicationInputProcessingConfigurationCmd.Flags().String("current-application-version-id", "", "Version of the application to which you want to add the input processing configuration.")
	kinesisanalytics_addApplicationInputProcessingConfigurationCmd.Flags().String("input-id", "", "The ID of the input configuration to add the input processing configuration to.")
	kinesisanalytics_addApplicationInputProcessingConfigurationCmd.Flags().String("input-processing-configuration", "", "The [InputProcessingConfiguration](https://docs.aws.amazon.com/kinesisanalytics/latest/dev/API_InputProcessingConfiguration.html) to add to the application.")
	kinesisanalytics_addApplicationInputProcessingConfigurationCmd.MarkFlagRequired("application-name")
	kinesisanalytics_addApplicationInputProcessingConfigurationCmd.MarkFlagRequired("current-application-version-id")
	kinesisanalytics_addApplicationInputProcessingConfigurationCmd.MarkFlagRequired("input-id")
	kinesisanalytics_addApplicationInputProcessingConfigurationCmd.MarkFlagRequired("input-processing-configuration")
	kinesisanalyticsCmd.AddCommand(kinesisanalytics_addApplicationInputProcessingConfigurationCmd)
}
