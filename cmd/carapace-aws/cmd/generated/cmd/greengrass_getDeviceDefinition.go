package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getDeviceDefinitionCmd = &cobra.Command{
	Use:   "get-device-definition",
	Short: "Retrieves information about a device definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getDeviceDefinitionCmd).Standalone()

	greengrass_getDeviceDefinitionCmd.Flags().String("device-definition-id", "", "The ID of the device definition.")
	greengrass_getDeviceDefinitionCmd.MarkFlagRequired("device-definition-id")
	greengrassCmd.AddCommand(greengrass_getDeviceDefinitionCmd)
}
