package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_enableDomainAutoRenewCmd = &cobra.Command{
	Use:   "enable-domain-auto-renew",
	Short: "This operation configures Amazon Route 53 to automatically renew the specified domain before the domain registration expires.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_enableDomainAutoRenewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_enableDomainAutoRenewCmd).Standalone()

		route53domains_enableDomainAutoRenewCmd.Flags().String("domain-name", "", "The name of the domain that you want to enable automatic renewal for.")
		route53domains_enableDomainAutoRenewCmd.MarkFlagRequired("domain-name")
	})
	route53domainsCmd.AddCommand(route53domains_enableDomainAutoRenewCmd)
}
