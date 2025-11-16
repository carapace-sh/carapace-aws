package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_addApplicationInputProcessingConfigurationCmd = &cobra.Command{
	Use:   "add-application-input-processing-configuration",
	Short: "Adds an [InputProcessingConfiguration]() to a SQL-based Kinesis Data Analytics application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_addApplicationInputProcessingConfigurationCmd).Standalone()

	kinesisanalyticsv2_addApplicationInputProcessingConfigurationCmd.Flags().String("application-name", "", "The name of the application to which you want to add the input processing configuration.")
	kinesisanalyticsv2_addApplicationInputProcessingConfigurationCmd.Flags().String("current-application-version-id", "", "The version of the application to which you want to add the input processing configuration.")
	kinesisanalyticsv2_addApplicationInputProcessingConfigurationCmd.Flags().String("input-id", "", "The ID of the input configuration to add the input processing configuration to.")
	kinesisanalyticsv2_addApplicationInputProcessingConfigurationCmd.Flags().String("input-processing-configuration", "", "The [InputProcessingConfiguration]() to add to the application.")
	kinesisanalyticsv2_addApplicationInputProcessingConfigurationCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2_addApplicationInputProcessingConfigurationCmd.MarkFlagRequired("current-application-version-id")
	kinesisanalyticsv2_addApplicationInputProcessingConfigurationCmd.MarkFlagRequired("input-id")
	kinesisanalyticsv2_addApplicationInputProcessingConfigurationCmd.MarkFlagRequired("input-processing-configuration")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_addApplicationInputProcessingConfigurationCmd)
}
