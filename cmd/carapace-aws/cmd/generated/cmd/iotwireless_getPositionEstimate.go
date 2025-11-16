package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getPositionEstimateCmd = &cobra.Command{
	Use:   "get-position-estimate",
	Short: "Get estimated position information as a payload in GeoJSON format.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getPositionEstimateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_getPositionEstimateCmd).Standalone()

		iotwireless_getPositionEstimateCmd.Flags().String("cell-towers", "", "Retrieves an estimated device position by resolving measurement data from cellular radio towers.")
		iotwireless_getPositionEstimateCmd.Flags().String("gnss", "", "Retrieves an estimated device position by resolving the global navigation satellite system (GNSS) scan data.")
		iotwireless_getPositionEstimateCmd.Flags().String("ip", "", "Retrieves an estimated device position by resolving the IP address information from the device.")
		iotwireless_getPositionEstimateCmd.Flags().String("timestamp", "", "Optional information that specifies the time when the position information will be resolved.")
		iotwireless_getPositionEstimateCmd.Flags().String("wi-fi-access-points", "", "Retrieves an estimated device position by resolving WLAN measurement data.")
	})
	iotwirelessCmd.AddCommand(iotwireless_getPositionEstimateCmd)
}
