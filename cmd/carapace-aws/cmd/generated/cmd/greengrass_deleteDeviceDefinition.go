package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_deleteDeviceDefinitionCmd = &cobra.Command{
	Use:   "delete-device-definition",
	Short: "Deletes a device definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_deleteDeviceDefinitionCmd).Standalone()

	greengrass_deleteDeviceDefinitionCmd.Flags().String("device-definition-id", "", "The ID of the device definition.")
	greengrass_deleteDeviceDefinitionCmd.MarkFlagRequired("device-definition-id")
	greengrassCmd.AddCommand(greengrass_deleteDeviceDefinitionCmd)
}
