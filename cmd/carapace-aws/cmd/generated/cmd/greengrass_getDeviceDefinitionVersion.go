package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_getDeviceDefinitionVersionCmd = &cobra.Command{
	Use:   "get-device-definition-version",
	Short: "Retrieves information about a device definition version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_getDeviceDefinitionVersionCmd).Standalone()

	greengrass_getDeviceDefinitionVersionCmd.Flags().String("device-definition-id", "", "The ID of the device definition.")
	greengrass_getDeviceDefinitionVersionCmd.Flags().String("device-definition-version-id", "", "The ID of the device definition version.")
	greengrass_getDeviceDefinitionVersionCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	greengrass_getDeviceDefinitionVersionCmd.MarkFlagRequired("device-definition-id")
	greengrass_getDeviceDefinitionVersionCmd.MarkFlagRequired("device-definition-version-id")
	greengrassCmd.AddCommand(greengrass_getDeviceDefinitionVersionCmd)
}
