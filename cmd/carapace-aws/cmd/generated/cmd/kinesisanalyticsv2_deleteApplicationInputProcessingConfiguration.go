package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_deleteApplicationInputProcessingConfigurationCmd = &cobra.Command{
	Use:   "delete-application-input-processing-configuration",
	Short: "Deletes an [InputProcessingConfiguration]() from an input.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_deleteApplicationInputProcessingConfigurationCmd).Standalone()

	kinesisanalyticsv2_deleteApplicationInputProcessingConfigurationCmd.Flags().String("application-name", "", "The name of the application.")
	kinesisanalyticsv2_deleteApplicationInputProcessingConfigurationCmd.Flags().String("current-application-version-id", "", "The application version.")
	kinesisanalyticsv2_deleteApplicationInputProcessingConfigurationCmd.Flags().String("input-id", "", "The ID of the input configuration from which to delete the input processing configuration.")
	kinesisanalyticsv2_deleteApplicationInputProcessingConfigurationCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2_deleteApplicationInputProcessingConfigurationCmd.MarkFlagRequired("current-application-version-id")
	kinesisanalyticsv2_deleteApplicationInputProcessingConfigurationCmd.MarkFlagRequired("input-id")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_deleteApplicationInputProcessingConfigurationCmd)
}
