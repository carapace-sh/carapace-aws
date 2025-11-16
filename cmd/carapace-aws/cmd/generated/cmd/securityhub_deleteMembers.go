package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_deleteMembersCmd = &cobra.Command{
	Use:   "delete-members",
	Short: "Deletes the specified member accounts from Security Hub.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_deleteMembersCmd).Standalone()

	securityhub_deleteMembersCmd.Flags().String("account-ids", "", "The list of account IDs for the member accounts to delete.")
	securityhub_deleteMembersCmd.MarkFlagRequired("account-ids")
	securityhubCmd.AddCommand(securityhub_deleteMembersCmd)
}
