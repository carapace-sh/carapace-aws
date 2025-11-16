package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_getImpersonationRoleCmd = &cobra.Command{
	Use:   "get-impersonation-role",
	Short: "Gets the impersonation role details for the given WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_getImpersonationRoleCmd).Standalone()

	workmail_getImpersonationRoleCmd.Flags().String("impersonation-role-id", "", "The impersonation role ID to retrieve.")
	workmail_getImpersonationRoleCmd.Flags().String("organization-id", "", "The WorkMail organization from which to retrieve the impersonation role.")
	workmail_getImpersonationRoleCmd.MarkFlagRequired("impersonation-role-id")
	workmail_getImpersonationRoleCmd.MarkFlagRequired("organization-id")
	workmailCmd.AddCommand(workmail_getImpersonationRoleCmd)
}
