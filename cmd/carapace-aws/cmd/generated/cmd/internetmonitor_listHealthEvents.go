package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_listHealthEventsCmd = &cobra.Command{
	Use:   "list-health-events",
	Short: "Lists all health events for a monitor in Amazon CloudWatch Internet Monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_listHealthEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(internetmonitor_listHealthEventsCmd).Standalone()

		internetmonitor_listHealthEventsCmd.Flags().String("end-time", "", "The time when a health event ended.")
		internetmonitor_listHealthEventsCmd.Flags().String("event-status", "", "The status of a health event.")
		internetmonitor_listHealthEventsCmd.Flags().String("linked-account-id", "", "The account ID for an account that you've set up cross-account sharing for in Amazon CloudWatch Internet Monitor.")
		internetmonitor_listHealthEventsCmd.Flags().String("max-results", "", "The number of health event objects that you want to return with this call.")
		internetmonitor_listHealthEventsCmd.Flags().String("monitor-name", "", "The name of the monitor.")
		internetmonitor_listHealthEventsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		internetmonitor_listHealthEventsCmd.Flags().String("start-time", "", "The time when a health event started.")
		internetmonitor_listHealthEventsCmd.MarkFlagRequired("monitor-name")
	})
	internetmonitorCmd.AddCommand(internetmonitor_listHealthEventsCmd)
}
