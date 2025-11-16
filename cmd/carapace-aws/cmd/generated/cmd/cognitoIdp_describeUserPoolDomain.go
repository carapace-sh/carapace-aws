package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_describeUserPoolDomainCmd = &cobra.Command{
	Use:   "describe-user-pool-domain",
	Short: "Given a user pool domain name, returns information about the domain configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_describeUserPoolDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_describeUserPoolDomainCmd).Standalone()

		cognitoIdp_describeUserPoolDomainCmd.Flags().String("domain", "", "The domain that you want to describe.")
		cognitoIdp_describeUserPoolDomainCmd.MarkFlagRequired("domain")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_describeUserPoolDomainCmd)
}
