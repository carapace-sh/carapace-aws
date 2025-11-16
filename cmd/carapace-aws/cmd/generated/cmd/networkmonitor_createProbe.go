package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmonitor_createProbeCmd = &cobra.Command{
	Use:   "create-probe",
	Short: "Create a probe within a monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmonitor_createProbeCmd).Standalone()

	networkmonitor_createProbeCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier to ensure the idempotency of the request.")
	networkmonitor_createProbeCmd.Flags().String("monitor-name", "", "The name of the monitor to associated with the probe.")
	networkmonitor_createProbeCmd.Flags().String("probe", "", "Describes the details of an individual probe for a monitor.")
	networkmonitor_createProbeCmd.Flags().String("tags", "", "The list of key-value pairs created and assigned to the probe.")
	networkmonitor_createProbeCmd.MarkFlagRequired("monitor-name")
	networkmonitor_createProbeCmd.MarkFlagRequired("probe")
	networkmonitorCmd.AddCommand(networkmonitor_createProbeCmd)
}
