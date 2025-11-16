package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_getDeviceCmd = &cobra.Command{
	Use:   "get-device",
	Short: "Retrieves the devices available in Amazon Braket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_getDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(braket_getDeviceCmd).Standalone()

		braket_getDeviceCmd.Flags().String("device-arn", "", "The ARN of the device to retrieve.")
		braket_getDeviceCmd.MarkFlagRequired("device-arn")
	})
	braketCmd.AddCommand(braket_getDeviceCmd)
}
