package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_putApplicationAuthenticationMethodCmd = &cobra.Command{
	Use:   "put-application-authentication-method",
	Short: "Adds or updates an authentication method for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_putApplicationAuthenticationMethodCmd).Standalone()

	ssoAdmin_putApplicationAuthenticationMethodCmd.Flags().String("application-arn", "", "Specifies the ARN of the application with the authentication method to add or update.")
	ssoAdmin_putApplicationAuthenticationMethodCmd.Flags().String("authentication-method", "", "Specifies a structure that describes the authentication method to add or update.")
	ssoAdmin_putApplicationAuthenticationMethodCmd.Flags().String("authentication-method-type", "", "Specifies the type of the authentication method that you want to add or update.")
	ssoAdmin_putApplicationAuthenticationMethodCmd.MarkFlagRequired("application-arn")
	ssoAdmin_putApplicationAuthenticationMethodCmd.MarkFlagRequired("authentication-method")
	ssoAdmin_putApplicationAuthenticationMethodCmd.MarkFlagRequired("authentication-method-type")
	ssoAdminCmd.AddCommand(ssoAdmin_putApplicationAuthenticationMethodCmd)
}
