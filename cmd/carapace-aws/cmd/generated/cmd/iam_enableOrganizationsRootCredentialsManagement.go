package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_enableOrganizationsRootCredentialsManagementCmd = &cobra.Command{
	Use:   "enable-organizations-root-credentials-management",
	Short: "Enables the management of privileged root user credentials across member accounts in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_enableOrganizationsRootCredentialsManagementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_enableOrganizationsRootCredentialsManagementCmd).Standalone()

	})
	iamCmd.AddCommand(iam_enableOrganizationsRootCredentialsManagementCmd)
}
