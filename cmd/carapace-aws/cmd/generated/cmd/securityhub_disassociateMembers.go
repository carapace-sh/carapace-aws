package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_disassociateMembersCmd = &cobra.Command{
	Use:   "disassociate-members",
	Short: "Disassociates the specified member accounts from the associated administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_disassociateMembersCmd).Standalone()

	securityhub_disassociateMembersCmd.Flags().String("account-ids", "", "The account IDs of the member accounts to disassociate from the administrator account.")
	securityhub_disassociateMembersCmd.MarkFlagRequired("account-ids")
	securityhubCmd.AddCommand(securityhub_disassociateMembersCmd)
}
