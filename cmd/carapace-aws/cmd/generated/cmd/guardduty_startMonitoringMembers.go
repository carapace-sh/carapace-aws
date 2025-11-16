package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_startMonitoringMembersCmd = &cobra.Command{
	Use:   "start-monitoring-members",
	Short: "Turns on GuardDuty monitoring of the specified member accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_startMonitoringMembersCmd).Standalone()

	guardduty_startMonitoringMembersCmd.Flags().String("account-ids", "", "A list of account IDs of the GuardDuty member accounts to start monitoring.")
	guardduty_startMonitoringMembersCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty administrator account associated with the member accounts to monitor.")
	guardduty_startMonitoringMembersCmd.MarkFlagRequired("account-ids")
	guardduty_startMonitoringMembersCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_startMonitoringMembersCmd)
}
