package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_disableOrganizationsRootCredentialsManagementCmd = &cobra.Command{
	Use:   "disable-organizations-root-credentials-management",
	Short: "Disables the management of privileged root user credentials across member accounts in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_disableOrganizationsRootCredentialsManagementCmd).Standalone()

	iamCmd.AddCommand(iam_disableOrganizationsRootCredentialsManagementCmd)
}
