package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_createUserPoolDomainCmd = &cobra.Command{
	Use:   "create-user-pool-domain",
	Short: "A user pool domain hosts managed login, an authorization server and web server for authentication in your application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_createUserPoolDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_createUserPoolDomainCmd).Standalone()

		cognitoIdp_createUserPoolDomainCmd.Flags().String("custom-domain-config", "", "The configuration for a custom domain.")
		cognitoIdp_createUserPoolDomainCmd.Flags().String("domain", "", "The domain string.")
		cognitoIdp_createUserPoolDomainCmd.Flags().String("managed-login-version", "", "The version of managed login branding that you want to apply to your domain.")
		cognitoIdp_createUserPoolDomainCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to add a domain.")
		cognitoIdp_createUserPoolDomainCmd.MarkFlagRequired("domain")
		cognitoIdp_createUserPoolDomainCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_createUserPoolDomainCmd)
}
