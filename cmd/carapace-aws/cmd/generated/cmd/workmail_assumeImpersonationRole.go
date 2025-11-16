package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_assumeImpersonationRoleCmd = &cobra.Command{
	Use:   "assume-impersonation-role",
	Short: "Assumes an impersonation role for the given WorkMail organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_assumeImpersonationRoleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_assumeImpersonationRoleCmd).Standalone()

		workmail_assumeImpersonationRoleCmd.Flags().String("impersonation-role-id", "", "The impersonation role ID to assume.")
		workmail_assumeImpersonationRoleCmd.Flags().String("organization-id", "", "The WorkMail organization under which the impersonation role will be assumed.")
		workmail_assumeImpersonationRoleCmd.MarkFlagRequired("impersonation-role-id")
		workmail_assumeImpersonationRoleCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_assumeImpersonationRoleCmd)
}
