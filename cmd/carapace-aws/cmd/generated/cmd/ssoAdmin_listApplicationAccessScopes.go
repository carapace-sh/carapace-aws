package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listApplicationAccessScopesCmd = &cobra.Command{
	Use:   "list-application-access-scopes",
	Short: "Lists the access scopes and authorized targets associated with an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listApplicationAccessScopesCmd).Standalone()

	ssoAdmin_listApplicationAccessScopesCmd.Flags().String("application-arn", "", "Specifies the ARN of the application.")
	ssoAdmin_listApplicationAccessScopesCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included in each response.")
	ssoAdmin_listApplicationAccessScopesCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	ssoAdmin_listApplicationAccessScopesCmd.MarkFlagRequired("application-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_listApplicationAccessScopesCmd)
}
