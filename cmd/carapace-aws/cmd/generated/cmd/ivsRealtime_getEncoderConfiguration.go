package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_getEncoderConfigurationCmd = &cobra.Command{
	Use:   "get-encoder-configuration",
	Short: "Gets information about the specified EncoderConfiguration resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_getEncoderConfigurationCmd).Standalone()

	ivsRealtime_getEncoderConfigurationCmd.Flags().String("arn", "", "ARN of the EncoderConfiguration resource.")
	ivsRealtime_getEncoderConfigurationCmd.MarkFlagRequired("arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_getEncoderConfigurationCmd)
}
