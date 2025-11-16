package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_deleteMembersCmd = &cobra.Command{
	Use:   "delete-members",
	Short: "Deletes GuardDuty member accounts (to the current GuardDuty administrator account) specified by the account IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_deleteMembersCmd).Standalone()

	guardduty_deleteMembersCmd.Flags().String("account-ids", "", "A list of account IDs of the GuardDuty member accounts that you want to delete.")
	guardduty_deleteMembersCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty account whose members you want to delete.")
	guardduty_deleteMembersCmd.MarkFlagRequired("account-ids")
	guardduty_deleteMembersCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_deleteMembersCmd)
}
