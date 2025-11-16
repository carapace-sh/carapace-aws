package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitor_deleteProbeCmd = &cobra.Command{
	Use:   "delete-probe",
	Short: "Deletes the specified probe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitor_deleteProbeCmd).Standalone()

	networkmonitor_deleteProbeCmd.Flags().String("monitor-name", "", "The name of the monitor to delete.")
	networkmonitor_deleteProbeCmd.Flags().String("probe-id", "", "The ID of the probe to delete.")
	networkmonitor_deleteProbeCmd.MarkFlagRequired("monitor-name")
	networkmonitor_deleteProbeCmd.MarkFlagRequired("probe-id")
	networkmonitorCmd.AddCommand(networkmonitor_deleteProbeCmd)
}
