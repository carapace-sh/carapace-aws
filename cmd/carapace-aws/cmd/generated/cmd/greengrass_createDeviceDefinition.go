package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_createDeviceDefinitionCmd = &cobra.Command{
	Use:   "create-device-definition",
	Short: "Creates a device definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_createDeviceDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrass_createDeviceDefinitionCmd).Standalone()

		greengrass_createDeviceDefinitionCmd.Flags().String("amzn-client-token", "", "A client token used to correlate requests and responses.")
		greengrass_createDeviceDefinitionCmd.Flags().String("initial-version", "", "Information about the initial version of the device definition.")
		greengrass_createDeviceDefinitionCmd.Flags().String("name", "", "The name of the device definition.")
		greengrass_createDeviceDefinitionCmd.Flags().String("tags", "", "Tag(s) to add to the new resource.")
	})
	greengrassCmd.AddCommand(greengrass_createDeviceDefinitionCmd)
}
