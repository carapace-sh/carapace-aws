package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerEdge_getDeviceRegistrationCmd = &cobra.Command{
	Use:   "get-device-registration",
	Short: "Use to check if a device is registered with SageMaker Edge Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerEdge_getDeviceRegistrationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerEdge_getDeviceRegistrationCmd).Standalone()

		sagemakerEdge_getDeviceRegistrationCmd.Flags().String("device-fleet-name", "", "The name of the fleet that the device belongs to.")
		sagemakerEdge_getDeviceRegistrationCmd.Flags().String("device-name", "", "The unique name of the device you want to get the registration status from.")
		sagemakerEdge_getDeviceRegistrationCmd.MarkFlagRequired("device-fleet-name")
		sagemakerEdge_getDeviceRegistrationCmd.MarkFlagRequired("device-name")
	})
	sagemakerEdgeCmd.AddCommand(sagemakerEdge_getDeviceRegistrationCmd)
}
