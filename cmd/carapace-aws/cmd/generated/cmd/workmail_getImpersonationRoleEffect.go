package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_getImpersonationRoleEffectCmd = &cobra.Command{
	Use:   "get-impersonation-role-effect",
	Short: "Tests whether the given impersonation role can impersonate a target user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_getImpersonationRoleEffectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_getImpersonationRoleEffectCmd).Standalone()

		workmail_getImpersonationRoleEffectCmd.Flags().String("impersonation-role-id", "", "The impersonation role ID to test.")
		workmail_getImpersonationRoleEffectCmd.Flags().String("organization-id", "", "The WorkMail organization where the impersonation role is defined.")
		workmail_getImpersonationRoleEffectCmd.Flags().String("target-user", "", "The WorkMail organization user chosen to test the impersonation role.")
		workmail_getImpersonationRoleEffectCmd.MarkFlagRequired("impersonation-role-id")
		workmail_getImpersonationRoleEffectCmd.MarkFlagRequired("organization-id")
		workmail_getImpersonationRoleEffectCmd.MarkFlagRequired("target-user")
	})
	workmailCmd.AddCommand(workmail_getImpersonationRoleEffectCmd)
}
