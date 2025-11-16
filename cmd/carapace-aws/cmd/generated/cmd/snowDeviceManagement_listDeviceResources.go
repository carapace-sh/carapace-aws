package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowDeviceManagement_listDeviceResourcesCmd = &cobra.Command{
	Use:   "list-device-resources",
	Short: "Returns a list of the Amazon Web Services resources available for a device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowDeviceManagement_listDeviceResourcesCmd).Standalone()

	snowDeviceManagement_listDeviceResourcesCmd.Flags().String("managed-device-id", "", "The ID of the managed device that you are listing the resources of.")
	snowDeviceManagement_listDeviceResourcesCmd.Flags().String("max-results", "", "The maximum number of resources per page.")
	snowDeviceManagement_listDeviceResourcesCmd.Flags().String("next-token", "", "A pagination token to continue to the next page of results.")
	snowDeviceManagement_listDeviceResourcesCmd.Flags().String("type", "", "A structure used to filter the results by type of resource.")
	snowDeviceManagement_listDeviceResourcesCmd.MarkFlagRequired("managed-device-id")
	snowDeviceManagementCmd.AddCommand(snowDeviceManagement_listDeviceResourcesCmd)
}
