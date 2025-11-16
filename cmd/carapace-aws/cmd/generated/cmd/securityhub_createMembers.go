package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_createMembersCmd = &cobra.Command{
	Use:   "create-members",
	Short: "Creates a member association in Security Hub between the specified accounts and the account used to make the request, which is the administrator account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_createMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_createMembersCmd).Standalone()

		securityhub_createMembersCmd.Flags().String("account-details", "", "The list of accounts to associate with the Security Hub administrator account.")
		securityhub_createMembersCmd.MarkFlagRequired("account-details")
	})
	securityhubCmd.AddCommand(securityhub_createMembersCmd)
}
