package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_listApplicationAuthenticationMethodsCmd = &cobra.Command{
	Use:   "list-application-authentication-methods",
	Short: "Lists all of the authentication methods supported by the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_listApplicationAuthenticationMethodsCmd).Standalone()

	ssoAdmin_listApplicationAuthenticationMethodsCmd.Flags().String("application-arn", "", "Specifies the ARN of the application with the authentication methods you want to list.")
	ssoAdmin_listApplicationAuthenticationMethodsCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
	ssoAdmin_listApplicationAuthenticationMethodsCmd.MarkFlagRequired("application-arn")
	ssoAdminCmd.AddCommand(ssoAdmin_listApplicationAuthenticationMethodsCmd)
}
