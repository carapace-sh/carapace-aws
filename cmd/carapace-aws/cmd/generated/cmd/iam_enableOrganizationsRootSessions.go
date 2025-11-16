package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_enableOrganizationsRootSessionsCmd = &cobra.Command{
	Use:   "enable-organizations-root-sessions",
	Short: "Allows the management account or delegated administrator to perform privileged tasks on member accounts in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_enableOrganizationsRootSessionsCmd).Standalone()

	iamCmd.AddCommand(iam_enableOrganizationsRootSessionsCmd)
}
