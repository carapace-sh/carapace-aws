package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_inviteAccountToOrganizationCmd = &cobra.Command{
	Use:   "invite-account-to-organization",
	Short: "Sends an invitation to another account to join your organization as a member account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_inviteAccountToOrganizationCmd).Standalone()

	organizations_inviteAccountToOrganizationCmd.Flags().String("notes", "", "Additional information that you want to include in the generated email to the recipient account owner.")
	organizations_inviteAccountToOrganizationCmd.Flags().String("tags", "", "A list of tags that you want to attach to the account when it becomes a member of the organization.")
	organizations_inviteAccountToOrganizationCmd.Flags().String("target", "", "The identifier (ID) of the Amazon Web Services account that you want to invite to join your organization.")
	organizations_inviteAccountToOrganizationCmd.MarkFlagRequired("target")
	organizationsCmd.AddCommand(organizations_inviteAccountToOrganizationCmd)
}
