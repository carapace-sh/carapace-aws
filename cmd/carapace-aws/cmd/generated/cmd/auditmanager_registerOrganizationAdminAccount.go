package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_registerOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "register-organization-admin-account",
	Short: "Enables an Amazon Web Services account within the organization as the delegated administrator for Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_registerOrganizationAdminAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_registerOrganizationAdminAccountCmd).Standalone()

		auditmanager_registerOrganizationAdminAccountCmd.Flags().String("admin-account-id", "", "The identifier for the delegated administrator account.")
		auditmanager_registerOrganizationAdminAccountCmd.MarkFlagRequired("admin-account-id")
	})
	auditmanagerCmd.AddCommand(auditmanager_registerOrganizationAdminAccountCmd)
}
