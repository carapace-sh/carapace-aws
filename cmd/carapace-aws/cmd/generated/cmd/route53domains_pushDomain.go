package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_pushDomainCmd = &cobra.Command{
	Use:   "push-domain",
	Short: "Moves a domain from Amazon Web Services to another registrar.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_pushDomainCmd).Standalone()

	route53domains_pushDomainCmd.Flags().String("domain-name", "", "Name of the domain.")
	route53domains_pushDomainCmd.Flags().String("target", "", "New IPS tag for the domain.")
	route53domains_pushDomainCmd.MarkFlagRequired("domain-name")
	route53domains_pushDomainCmd.MarkFlagRequired("target")
	route53domainsCmd.AddCommand(route53domains_pushDomainCmd)
}
