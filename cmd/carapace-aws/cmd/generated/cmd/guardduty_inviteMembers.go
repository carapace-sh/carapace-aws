package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_inviteMembersCmd = &cobra.Command{
	Use:   "invite-members",
	Short: "Invites Amazon Web Services accounts to become members of an organization administered by the Amazon Web Services account that invokes this API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_inviteMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_inviteMembersCmd).Standalone()

		guardduty_inviteMembersCmd.Flags().String("account-ids", "", "A list of account IDs of the accounts that you want to invite to GuardDuty as members.")
		guardduty_inviteMembersCmd.Flags().String("detector-id", "", "The unique ID of the detector of the GuardDuty account with which you want to invite members.")
		guardduty_inviteMembersCmd.Flags().Bool("disable-email-notification", false, "A Boolean value that specifies whether you want to disable email notification to the accounts that you are inviting to GuardDuty as members.")
		guardduty_inviteMembersCmd.Flags().String("message", "", "The invitation message that you want to send to the accounts that you're inviting to GuardDuty as members.")
		guardduty_inviteMembersCmd.Flags().Bool("no-disable-email-notification", false, "A Boolean value that specifies whether you want to disable email notification to the accounts that you are inviting to GuardDuty as members.")
		guardduty_inviteMembersCmd.MarkFlagRequired("account-ids")
		guardduty_inviteMembersCmd.MarkFlagRequired("detector-id")
		guardduty_inviteMembersCmd.Flag("no-disable-email-notification").Hidden = true
	})
	guarddutyCmd.AddCommand(guardduty_inviteMembersCmd)
}
