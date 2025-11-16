package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_createImpersonationRoleCmd = &cobra.Command{
	Use:   "create-impersonation-role",
	Short: "Creates an impersonation role for the given WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_createImpersonationRoleCmd).Standalone()

	workmail_createImpersonationRoleCmd.Flags().String("client-token", "", "The idempotency token for the client request.")
	workmail_createImpersonationRoleCmd.Flags().String("description", "", "The description of the new impersonation role.")
	workmail_createImpersonationRoleCmd.Flags().String("name", "", "The name of the new impersonation role.")
	workmail_createImpersonationRoleCmd.Flags().String("organization-id", "", "The WorkMail organization to create the new impersonation role within.")
	workmail_createImpersonationRoleCmd.Flags().String("rules", "", "The list of rules for the impersonation role.")
	workmail_createImpersonationRoleCmd.Flags().String("type", "", "The impersonation role's type.")
	workmail_createImpersonationRoleCmd.MarkFlagRequired("name")
	workmail_createImpersonationRoleCmd.MarkFlagRequired("organization-id")
	workmail_createImpersonationRoleCmd.MarkFlagRequired("rules")
	workmail_createImpersonationRoleCmd.MarkFlagRequired("type")
	workmailCmd.AddCommand(workmail_createImpersonationRoleCmd)
}
