package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_deleteImpersonationRoleCmd = &cobra.Command{
	Use:   "delete-impersonation-role",
	Short: "Deletes an impersonation role for the given WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_deleteImpersonationRoleCmd).Standalone()

	workmail_deleteImpersonationRoleCmd.Flags().String("impersonation-role-id", "", "The ID of the impersonation role to delete.")
	workmail_deleteImpersonationRoleCmd.Flags().String("organization-id", "", "The WorkMail organization from which to delete the impersonation role.")
	workmail_deleteImpersonationRoleCmd.MarkFlagRequired("impersonation-role-id")
	workmail_deleteImpersonationRoleCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_deleteImpersonationRoleCmd)
}
