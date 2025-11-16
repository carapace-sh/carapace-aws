package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_disassociateMembersCmd = &cobra.Command{
	Use:   "disassociate-members",
	Short: "Disassociates GuardDuty member accounts (from the current administrator account) specified by the account IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_disassociateMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_disassociateMembersCmd).Standalone()

		guardduty_disassociateMembersCmd.Flags().String("account-ids", "", "A list of account IDs of the GuardDuty member accounts that you want to disassociate from the administrator account.")
		guardduty_disassociateMembersCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty account whose members you want to disassociate from the administrator account.")
		guardduty_disassociateMembersCmd.MarkFlagRequired("account-ids")
		guardduty_disassociateMembersCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_disassociateMembersCmd)
}
