package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssoAdmin_deleteApplicationAuthenticationMethodCmd = &cobra.Command{
	Use:   "delete-application-authentication-method",
	Short: "Deletes an authentication method from an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssoAdmin_deleteApplicationAuthenticationMethodCmd).Standalone()

	ssoAdmin_deleteApplicationAuthenticationMethodCmd.Flags().String("application-arn", "", "Specifies the ARN of the application with the authentication method to delete.")
	ssoAdmin_deleteApplicationAuthenticationMethodCmd.Flags().String("authentication-method-type", "", "Specifies the authentication method type to delete from the application.")
	ssoAdmin_deleteApplicationAuthenticationMethodCmd.MarkFlagRequired("application-arn")
	ssoAdmin_deleteApplicationAuthenticationMethodCmd.MarkFlagRequired("authentication-method-type")
	ssoAdminCmd.AddCommand(ssoAdmin_deleteApplicationAuthenticationMethodCmd)
}
