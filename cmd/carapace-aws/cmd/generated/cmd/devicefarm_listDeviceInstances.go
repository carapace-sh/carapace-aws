package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_listDeviceInstancesCmd = &cobra.Command{
	Use:   "list-device-instances",
	Short: "Returns information about the private device instances associated with one or more AWS accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_listDeviceInstancesCmd).Standalone()

	devicefarm_listDeviceInstancesCmd.Flags().String("max-results", "", "An integer that specifies the maximum number of items you want to return in the API response.")
	devicefarm_listDeviceInstancesCmd.Flags().String("next-token", "", "An identifier that was returned from the previous call to this operation, which can be used to return the next set of items in the list.")
	devicefarmCmd.AddCommand(devicefarm_listDeviceInstancesCmd)
}
