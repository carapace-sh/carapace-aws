package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitor_getProbeCmd = &cobra.Command{
	Use:   "get-probe",
	Short: "Returns the details about a probe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitor_getProbeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmonitor_getProbeCmd).Standalone()

		networkmonitor_getProbeCmd.Flags().String("monitor-name", "", "The name of the monitor associated with the probe.")
		networkmonitor_getProbeCmd.Flags().String("probe-id", "", "The ID of the probe to get information about.")
		networkmonitor_getProbeCmd.MarkFlagRequired("monitor-name")
		networkmonitor_getProbeCmd.MarkFlagRequired("probe-id")
	})
	networkmonitorCmd.AddCommand(networkmonitor_getProbeCmd)
}
