package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var internetmonitor_getHealthEventCmd = &cobra.Command{
	Use:   "get-health-event",
	Short: "Gets information that Amazon CloudWatch Internet Monitor has created and stored about a health event for a specified monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(internetmonitor_getHealthEventCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(internetmonitor_getHealthEventCmd).Standalone()

		internetmonitor_getHealthEventCmd.Flags().String("event-id", "", "The internally-generated identifier of a health event.")
		internetmonitor_getHealthEventCmd.Flags().String("linked-account-id", "", "The account ID for an account that you've set up cross-account sharing for in Amazon CloudWatch Internet Monitor.")
		internetmonitor_getHealthEventCmd.Flags().String("monitor-name", "", "The name of the monitor.")
		internetmonitor_getHealthEventCmd.MarkFlagRequired("event-id")
		internetmonitor_getHealthEventCmd.MarkFlagRequired("monitor-name")
	})
	internetmonitorCmd.AddCommand(internetmonitor_getHealthEventCmd)
}
