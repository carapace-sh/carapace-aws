package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_getApplicationAuthenticationMethodCmd = &cobra.Command{
	Use:   "get-application-authentication-method",
	Short: "Retrieves details about an authentication method used by an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_getApplicationAuthenticationMethodCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssoAdmin_getApplicationAuthenticationMethodCmd).Standalone()

		ssoAdmin_getApplicationAuthenticationMethodCmd.Flags().String("application-arn", "", "Specifies the ARN of the application.")
		ssoAdmin_getApplicationAuthenticationMethodCmd.Flags().String("authentication-method-type", "", "Specifies the type of authentication method for which you want details.")
		ssoAdmin_getApplicationAuthenticationMethodCmd.MarkFlagRequired("application-arn")
		ssoAdmin_getApplicationAuthenticationMethodCmd.MarkFlagRequired("authentication-method-type")
	})
	ssoAdminCmd.AddCommand(ssoAdmin_getApplicationAuthenticationMethodCmd)
}
