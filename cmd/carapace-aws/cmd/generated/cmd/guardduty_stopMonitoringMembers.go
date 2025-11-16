package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_stopMonitoringMembersCmd = &cobra.Command{
	Use:   "stop-monitoring-members",
	Short: "Stops GuardDuty monitoring for the specified member accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_stopMonitoringMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_stopMonitoringMembersCmd).Standalone()

		guardduty_stopMonitoringMembersCmd.Flags().String("account-ids", "", "A list of account IDs for the member accounts to stop monitoring.")
		guardduty_stopMonitoringMembersCmd.Flags().String("detector-id", "", "The unique ID of the detector associated with the GuardDuty administrator account that is monitoring member accounts.")
		guardduty_stopMonitoringMembersCmd.MarkFlagRequired("account-ids")
		guardduty_stopMonitoringMembersCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_stopMonitoringMembersCmd)
}
