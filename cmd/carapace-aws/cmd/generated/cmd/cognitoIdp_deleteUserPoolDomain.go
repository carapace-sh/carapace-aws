package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_deleteUserPoolDomainCmd = &cobra.Command{
	Use:   "delete-user-pool-domain",
	Short: "Given a user pool ID and domain identifier, deletes a user pool domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_deleteUserPoolDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_deleteUserPoolDomainCmd).Standalone()

		cognitoIdp_deleteUserPoolDomainCmd.Flags().String("domain", "", "The domain that you want to delete.")
		cognitoIdp_deleteUserPoolDomainCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to delete the domain.")
		cognitoIdp_deleteUserPoolDomainCmd.MarkFlagRequired("domain")
		cognitoIdp_deleteUserPoolDomainCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_deleteUserPoolDomainCmd)
}
