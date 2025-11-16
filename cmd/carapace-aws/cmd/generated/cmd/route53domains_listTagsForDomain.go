package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_listTagsForDomainCmd = &cobra.Command{
	Use:   "list-tags-for-domain",
	Short: "This operation returns all of the tags that are associated with the specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_listTagsForDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_listTagsForDomainCmd).Standalone()

		route53domains_listTagsForDomainCmd.Flags().String("domain-name", "", "The domain for which you want to get a list of tags.")
		route53domains_listTagsForDomainCmd.MarkFlagRequired("domain-name")
	})
	route53domainsCmd.AddCommand(route53domains_listTagsForDomainCmd)
}
