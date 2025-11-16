package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_deleteRecordingConfigurationCmd = &cobra.Command{
	Use:   "delete-recording-configuration",
	Short: "Deletes the recording configuration for the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_deleteRecordingConfigurationCmd).Standalone()

	ivs_deleteRecordingConfigurationCmd.Flags().String("arn", "", "ARN of the recording configuration to be deleted.")
	ivs_deleteRecordingConfigurationCmd.MarkFlagRequired("arn")
	ivsCmd.AddCommand(ivs_deleteRecordingConfigurationCmd)
}
