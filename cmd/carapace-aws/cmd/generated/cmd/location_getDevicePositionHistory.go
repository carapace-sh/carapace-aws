package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_getDevicePositionHistoryCmd = &cobra.Command{
	Use:   "get-device-position-history",
	Short: "Retrieves the device position history from a tracker resource within a specified range of time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_getDevicePositionHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_getDevicePositionHistoryCmd).Standalone()

		location_getDevicePositionHistoryCmd.Flags().String("device-id", "", "The device whose position history you want to retrieve.")
		location_getDevicePositionHistoryCmd.Flags().String("end-time-exclusive", "", "Specify the end time for the position history in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format: `YYYY-MM-DDThh:mm:ss.sssZ`.")
		location_getDevicePositionHistoryCmd.Flags().String("max-results", "", "An optional limit for the number of device positions returned in a single call.")
		location_getDevicePositionHistoryCmd.Flags().String("next-token", "", "The pagination token specifying which page of results to return in the response.")
		location_getDevicePositionHistoryCmd.Flags().String("start-time-inclusive", "", "Specify the start time for the position history in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format: `YYYY-MM-DDThh:mm:ss.sssZ`.")
		location_getDevicePositionHistoryCmd.Flags().String("tracker-name", "", "The tracker resource receiving the request for the device position history.")
		location_getDevicePositionHistoryCmd.MarkFlagRequired("device-id")
		location_getDevicePositionHistoryCmd.MarkFlagRequired("tracker-name")
	})
	locationCmd.AddCommand(location_getDevicePositionHistoryCmd)
}
