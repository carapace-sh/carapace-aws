package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_putApplicationAccessScopeCmd = &cobra.Command{
	Use:   "put-application-access-scope",
	Short: "Adds or updates the list of authorized targets for an IAM Identity Center access scope for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_putApplicationAccessScopeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_putApplicationAccessScopeCmd).Standalone()

		ssoAdmin_putApplicationAccessScopeCmd.Flags().String("application-arn", "", "Specifies the ARN of the application with the access scope with the targets to add or update.")
		ssoAdmin_putApplicationAccessScopeCmd.Flags().String("authorized-targets", "", "Specifies an array list of ARNs that represent the authorized targets for this access scope.")
		ssoAdmin_putApplicationAccessScopeCmd.Flags().String("scope", "", "Specifies the name of the access scope to be associated with the specified targets.")
		ssoAdmin_putApplicationAccessScopeCmd.MarkFlagRequired("application-arn")
		ssoAdmin_putApplicationAccessScopeCmd.MarkFlagRequired("scope")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_putApplicationAccessScopeCmd)
}
