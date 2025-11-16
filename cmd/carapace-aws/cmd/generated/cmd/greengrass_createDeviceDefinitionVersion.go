package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createDeviceDefinitionVersionCmd = &cobra.Command{
	Use:   "create-device-definition-version",
	Short: "Creates a version of a device definition that has already been defined.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createDeviceDefinitionVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_createDeviceDefinitionVersionCmd).Standalone()

		greengrass_createDeviceDefinitionVersionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
		greengrass_createDeviceDefinitionVersionCmd.Flags().String("device-definition-id", "", "The ID of the device definition.")
		greengrass_createDeviceDefinitionVersionCmd.Flags().String("devices", "", "A list of devices in the definition version.")
		greengrass_createDeviceDefinitionVersionCmd.MarkFlagRequired("device-definition-id")
	})
	greengrassCmd.AddCommand(greengrass_createDeviceDefinitionVersionCmd)
}
