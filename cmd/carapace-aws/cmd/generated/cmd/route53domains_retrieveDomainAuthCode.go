package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_retrieveDomainAuthCodeCmd = &cobra.Command{
	Use:   "retrieve-domain-auth-code",
	Short: "This operation returns the authorization code for the domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_retrieveDomainAuthCodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_retrieveDomainAuthCodeCmd).Standalone()

		route53domains_retrieveDomainAuthCodeCmd.Flags().String("domain-name", "", "The name of the domain that you want to get an authorization code for.")
		route53domains_retrieveDomainAuthCodeCmd.MarkFlagRequired("domain-name")
	})
	route53domainsCmd.AddCommand(route53domains_retrieveDomainAuthCodeCmd)
}
