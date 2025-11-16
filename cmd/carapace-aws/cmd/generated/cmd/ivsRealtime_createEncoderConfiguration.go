package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_createEncoderConfigurationCmd = &cobra.Command{
	Use:   "create-encoder-configuration",
	Short: "Creates an EncoderConfiguration object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_createEncoderConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_createEncoderConfigurationCmd).Standalone()

		ivsRealtime_createEncoderConfigurationCmd.Flags().String("name", "", "Optional name to identify the resource.")
		ivsRealtime_createEncoderConfigurationCmd.Flags().String("tags", "", "Tags attached to the resource.")
		ivsRealtime_createEncoderConfigurationCmd.Flags().String("video", "", "Video configuration.")
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_createEncoderConfigurationCmd)
}
