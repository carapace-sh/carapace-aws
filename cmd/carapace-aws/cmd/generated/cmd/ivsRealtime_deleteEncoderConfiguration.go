package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_deleteEncoderConfigurationCmd = &cobra.Command{
	Use:   "delete-encoder-configuration",
	Short: "Deletes an EncoderConfiguration resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_deleteEncoderConfigurationCmd).Standalone()

	ivsRealtime_deleteEncoderConfigurationCmd.Flags().String("arn", "", "ARN of the EncoderConfiguration.")
	ivsRealtime_deleteEncoderConfigurationCmd.MarkFlagRequired("arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_deleteEncoderConfigurationCmd)
}
