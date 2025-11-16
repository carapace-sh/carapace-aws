package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteDeviceFleetCmd = &cobra.Command{
	Use:   "delete-device-fleet",
	Short: "Deletes a fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteDeviceFleetCmd).Standalone()

	sagemaker_deleteDeviceFleetCmd.Flags().String("device-fleet-name", "", "The name of the fleet to delete.")
	sagemaker_deleteDeviceFleetCmd.MarkFlagRequired("device-fleet-name")
	sagemakerCmd.AddCommand(sagemaker_deleteDeviceFleetCmd)
}
