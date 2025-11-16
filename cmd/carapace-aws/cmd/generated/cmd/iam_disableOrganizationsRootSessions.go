package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_disableOrganizationsRootSessionsCmd = &cobra.Command{
	Use:   "disable-organizations-root-sessions",
	Short: "Disables root user sessions for privileged tasks across member accounts in your organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_disableOrganizationsRootSessionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_disableOrganizationsRootSessionsCmd).Standalone()

	})
	iamCmd.AddCommand(iam_disableOrganizationsRootSessionsCmd)
}
