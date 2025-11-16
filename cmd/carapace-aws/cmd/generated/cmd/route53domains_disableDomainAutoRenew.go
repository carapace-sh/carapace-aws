package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_disableDomainAutoRenewCmd = &cobra.Command{
	Use:   "disable-domain-auto-renew",
	Short: "This operation disables automatic renewal of domain registration for the specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_disableDomainAutoRenewCmd).Standalone()

	route53domains_disableDomainAutoRenewCmd.Flags().String("domain-name", "", "The name of the domain that you want to disable automatic renewal for.")
	route53domains_disableDomainAutoRenewCmd.MarkFlagRequired("domain-name")
	route53domainsCmd.AddCommand(route53domains_disableDomainAutoRenewCmd)
}
