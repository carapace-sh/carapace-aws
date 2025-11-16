package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_deleteApplicationAccessScopeCmd = &cobra.Command{
	Use:   "delete-application-access-scope",
	Short: "Deletes an IAM Identity Center access scope from an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_deleteApplicationAccessScopeCmd).Standalone()

	ssoAdmin_deleteApplicationAccessScopeCmd.Flags().String("application-arn", "", "Specifies the ARN of the application with the access scope to delete.")
	ssoAdmin_deleteApplicationAccessScopeCmd.Flags().String("scope", "", "Specifies the name of the access scope to remove from the application.")
	ssoAdmin_deleteApplicationAccessScopeCmd.MarkFlagRequired("application-arn")
	ssoAdmin_deleteApplicationAccessScopeCmd.MarkFlagRequired("scope")
	ssoAdminCmd.AddCommand(ssoAdmin_deleteApplicationAccessScopeCmd)
}
