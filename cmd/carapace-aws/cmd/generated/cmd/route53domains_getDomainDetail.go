package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_getDomainDetailCmd = &cobra.Command{
	Use:   "get-domain-detail",
	Short: "This operation returns detailed information about a specified domain that is associated with the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_getDomainDetailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_getDomainDetailCmd).Standalone()

		route53domains_getDomainDetailCmd.Flags().String("domain-name", "", "The name of the domain that you want to get detailed information about.")
		route53domains_getDomainDetailCmd.MarkFlagRequired("domain-name")
	})
	route53domainsCmd.AddCommand(route53domains_getDomainDetailCmd)
}
