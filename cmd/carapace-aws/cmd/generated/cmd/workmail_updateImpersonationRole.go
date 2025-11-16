package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_updateImpersonationRoleCmd = &cobra.Command{
	Use:   "update-impersonation-role",
	Short: "Updates an impersonation role for the given WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_updateImpersonationRoleCmd).Standalone()

	workmail_updateImpersonationRoleCmd.Flags().String("description", "", "The updated impersonation role description.")
	workmail_updateImpersonationRoleCmd.Flags().String("impersonation-role-id", "", "The ID of the impersonation role to update.")
	workmail_updateImpersonationRoleCmd.Flags().String("name", "", "The updated impersonation role name.")
	workmail_updateImpersonationRoleCmd.Flags().String("organization-id", "", "The WorkMail organization that contains the impersonation role to update.")
	workmail_updateImpersonationRoleCmd.Flags().String("rules", "", "The updated list of rules.")
	workmail_updateImpersonationRoleCmd.Flags().String("type", "", "The updated impersonation role type.")
	workmail_updateImpersonationRoleCmd.MarkFlagRequired("impersonation-role-id")
	workmail_updateImpersonationRoleCmd.MarkFlagRequired("name")
	workmail_updateImpersonationRoleCmd.MarkFlagRequired("organization-id")
	workmail_updateImpersonationRoleCmd.MarkFlagRequired("rules")
	workmail_updateImpersonationRoleCmd.MarkFlagRequired("type")
	workmailCmd.AddCommand(workmail_updateImpersonationRoleCmd)
}
