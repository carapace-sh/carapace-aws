package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_createMembersCmd = &cobra.Command{
	Use:   "create-members",
	Short: "Creates member accounts of the current Amazon Web Services account by specifying a list of Amazon Web Services account IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_createMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_createMembersCmd).Standalone()

		guardduty_createMembersCmd.Flags().String("account-details", "", "A list of account ID and email address pairs of the accounts that you want to associate with the GuardDuty administrator account.")
		guardduty_createMembersCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty account for which you want to associate member accounts.")
		guardduty_createMembersCmd.MarkFlagRequired("account-details")
		guardduty_createMembersCmd.MarkFlagRequired("detector-id")
	})
	guarddutyCmd.AddCommand(guardduty_createMembersCmd)
}
