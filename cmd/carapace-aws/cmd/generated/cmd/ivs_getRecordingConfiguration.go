package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_getRecordingConfigurationCmd = &cobra.Command{
	Use:   "get-recording-configuration",
	Short: "Gets the recording configuration for the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_getRecordingConfigurationCmd).Standalone()

	ivs_getRecordingConfigurationCmd.Flags().String("arn", "", "ARN of the recording configuration to be retrieved.")
	ivs_getRecordingConfigurationCmd.MarkFlagRequired("arn")
	ivsCmd.AddCommand(ivs_getRecordingConfigurationCmd)
}
