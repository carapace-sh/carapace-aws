package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_describeCrossAccountAccessRoleCmd = &cobra.Command{
	Use:   "describe-cross-account-access-role",
	Short: "Describes the IAM role that enables Amazon Inspector to access your AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_describeCrossAccountAccessRoleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_describeCrossAccountAccessRoleCmd).Standalone()

	})
	inspectorCmd.AddCommand(inspector_describeCrossAccountAccessRoleCmd)
}
