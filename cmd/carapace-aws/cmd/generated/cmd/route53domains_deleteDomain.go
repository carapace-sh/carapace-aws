package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_deleteDomainCmd = &cobra.Command{
	Use:   "delete-domain",
	Short: "This operation deletes the specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_deleteDomainCmd).Standalone()

	route53domains_deleteDomainCmd.Flags().String("domain-name", "", "Name of the domain to be deleted.")
	route53domains_deleteDomainCmd.MarkFlagRequired("domain-name")
	route53domainsCmd.AddCommand(route53domains_deleteDomainCmd)
}
