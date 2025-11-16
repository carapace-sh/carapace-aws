package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getDeviceCmd = &cobra.Command{
	Use:   "get-device",
	Short: "Gets information about a unique device type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getDeviceCmd).Standalone()

	devicefarm_getDeviceCmd.Flags().String("arn", "", "The device type's ARN.")
	devicefarm_getDeviceCmd.MarkFlagRequired("arn")
	devicefarmCmd.AddCommand(devicefarm_getDeviceCmd)
}
