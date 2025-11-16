package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitor_updateProbeCmd = &cobra.Command{
	Use:   "update-probe",
	Short: "Updates a monitor probe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitor_updateProbeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmonitor_updateProbeCmd).Standalone()

		networkmonitor_updateProbeCmd.Flags().String("destination", "", "The updated IP address for the probe destination.")
		networkmonitor_updateProbeCmd.Flags().String("destination-port", "", "The updated port for the probe destination.")
		networkmonitor_updateProbeCmd.Flags().String("monitor-name", "", "The name of the monitor that the probe was updated for.")
		networkmonitor_updateProbeCmd.Flags().String("packet-size", "", "he updated packets size for network traffic between the source and destination.")
		networkmonitor_updateProbeCmd.Flags().String("probe-id", "", "The ID of the probe to update.")
		networkmonitor_updateProbeCmd.Flags().String("protocol", "", "The updated network protocol for the destination.")
		networkmonitor_updateProbeCmd.Flags().String("state", "", "The state of the probe update.")
		networkmonitor_updateProbeCmd.MarkFlagRequired("monitor-name")
		networkmonitor_updateProbeCmd.MarkFlagRequired("probe-id")
	})
	networkmonitorCmd.AddCommand(networkmonitor_updateProbeCmd)
}
