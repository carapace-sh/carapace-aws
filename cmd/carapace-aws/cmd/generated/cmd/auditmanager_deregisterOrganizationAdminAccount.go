package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_deregisterOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "deregister-organization-admin-account",
	Short: "Removes the specified Amazon Web Services account as a delegated administrator for Audit Manager.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_deregisterOrganizationAdminAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(auditmanager_deregisterOrganizationAdminAccountCmd).Standalone()

		auditmanager_deregisterOrganizationAdminAccountCmd.Flags().String("admin-account-id", "", "The identifier for the administrator account.")
	})
	auditmanagerCmd.AddCommand(auditmanager_deregisterOrganizationAdminAccountCmd)
}
