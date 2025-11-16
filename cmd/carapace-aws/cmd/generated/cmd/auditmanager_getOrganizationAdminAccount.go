package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var auditmanager_getOrganizationAdminAccountCmd = &cobra.Command{
	Use:   "get-organization-admin-account",
	Short: "Gets the name of the delegated Amazon Web Services administrator account for a specified organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auditmanager_getOrganizationAdminAccountCmd).Standalone()

	auditmanagerCmd.AddCommand(auditmanager_getOrganizationAdminAccountCmd)
}
