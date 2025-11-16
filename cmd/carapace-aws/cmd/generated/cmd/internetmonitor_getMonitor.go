package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_getMonitorCmd = &cobra.Command{
	Use:   "get-monitor",
	Short: "Gets information about a monitor in Amazon CloudWatch Internet Monitor based on a monitor name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_getMonitorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(internetmonitor_getMonitorCmd).Standalone()

		internetmonitor_getMonitorCmd.Flags().String("linked-account-id", "", "The account ID for an account that you've set up cross-account sharing for in Amazon CloudWatch Internet Monitor.")
		internetmonitor_getMonitorCmd.Flags().String("monitor-name", "", "The name of the monitor.")
		internetmonitor_getMonitorCmd.MarkFlagRequired("monitor-name")
	})
	internetmonitorCmd.AddCommand(internetmonitor_getMonitorCmd)
}
