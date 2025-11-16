package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_listDevicesCmd = &cobra.Command{
	Use:   "list-devices",
	Short: "Returns a list of devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_listDevicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_listDevicesCmd).Standalone()

		panorama_listDevicesCmd.Flags().String("device-aggregated-status-filter", "", "Filter based on a device's status.")
		panorama_listDevicesCmd.Flags().String("max-results", "", "The maximum number of devices to return in one page of results.")
		panorama_listDevicesCmd.Flags().String("name-filter", "", "Filter based on device's name.")
		panorama_listDevicesCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		panorama_listDevicesCmd.Flags().String("sort-by", "", "The target column to be sorted on.")
		panorama_listDevicesCmd.Flags().String("sort-order", "", "The sorting order for the returned list.")
	})
	panoramaCmd.AddCommand(panorama_listDevicesCmd)
}
