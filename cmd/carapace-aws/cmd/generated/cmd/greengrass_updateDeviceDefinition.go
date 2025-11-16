package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_updateDeviceDefinitionCmd = &cobra.Command{
	Use:   "update-device-definition",
	Short: "Updates a device definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_updateDeviceDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_updateDeviceDefinitionCmd).Standalone()

		greengrass_updateDeviceDefinitionCmd.Flags().String("device-definition-id", "", "The ID of the device definition.")
		greengrass_updateDeviceDefinitionCmd.Flags().String("name", "", "The name of the definition.")
		greengrass_updateDeviceDefinitionCmd.MarkFlagRequired("device-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_updateDeviceDefinitionCmd)
}
