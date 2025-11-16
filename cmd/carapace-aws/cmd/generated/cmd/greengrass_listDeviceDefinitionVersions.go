package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrass_listDeviceDefinitionVersionsCmd = &cobra.Command{
	Use:   "list-device-definition-versions",
	Short: "Lists the versions of a device definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrass_listDeviceDefinitionVersionsCmd).Standalone()

	greengrass_listDeviceDefinitionVersionsCmd.Flags().String("device-definition-id", "", "The ID of the device definition.")
	greengrass_listDeviceDefinitionVersionsCmd.Flags().String("max-results", "", "The maximum number of results to be returned per request.")
	greengrass_listDeviceDefinitionVersionsCmd.Flags().String("next-token", "", "The token for the next set of results, or ''null'' if there are no additional results.")
	greengrass_listDeviceDefinitionVersionsCmd.MarkFlagRequired("device-definition-id")
	greengrassCmd.AddCommand(greengrass_listDeviceDefinitionVersionsCmd)
}
