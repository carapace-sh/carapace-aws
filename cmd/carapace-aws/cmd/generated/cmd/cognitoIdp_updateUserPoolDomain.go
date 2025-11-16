package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_updateUserPoolDomainCmd = &cobra.Command{
	Use:   "update-user-pool-domain",
	Short: "A user pool domain hosts managed login, an authorization server and web server for authentication in your application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_updateUserPoolDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_updateUserPoolDomainCmd).Standalone()

		cognitoIdp_updateUserPoolDomainCmd.Flags().String("custom-domain-config", "", "The configuration for a custom domain that hosts managed login for your application.")
		cognitoIdp_updateUserPoolDomainCmd.Flags().String("domain", "", "The name of the domain that you want to update.")
		cognitoIdp_updateUserPoolDomainCmd.Flags().String("managed-login-version", "", "A version number that indicates the state of managed login for your domain.")
		cognitoIdp_updateUserPoolDomainCmd.Flags().String("user-pool-id", "", "The ID of the user pool that is associated with the domain you're updating.")
		cognitoIdp_updateUserPoolDomainCmd.MarkFlagRequired("domain")
		cognitoIdp_updateUserPoolDomainCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_updateUserPoolDomainCmd)
}
