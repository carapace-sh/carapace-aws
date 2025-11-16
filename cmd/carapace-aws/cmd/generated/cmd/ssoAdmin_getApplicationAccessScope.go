package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_getApplicationAccessScopeCmd = &cobra.Command{
	Use:   "get-application-access-scope",
	Short: "Retrieves the authorized targets for an IAM Identity Center access scope for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_getApplicationAccessScopeCmd).Standalone()

	ssoAdmin_getApplicationAccessScopeCmd.Flags().String("application-arn", "", "Specifies the ARN of the application with the access scope that you want to retrieve.")
	ssoAdmin_getApplicationAccessScopeCmd.Flags().String("scope", "", "Specifies the name of the access scope for which you want the authorized targets.")
	ssoAdmin_getApplicationAccessScopeCmd.MarkFlagRequired("application-arn")
	ssoAdmin_getApplicationAccessScopeCmd.MarkFlagRequired("scope")
	ssoAdminCmd.AddCommand(ssoAdmin_getApplicationAccessScopeCmd)
}
