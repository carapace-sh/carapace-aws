package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_deleteTagsForDomainCmd = &cobra.Command{
	Use:   "delete-tags-for-domain",
	Short: "This operation deletes the specified tags for a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_deleteTagsForDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_deleteTagsForDomainCmd).Standalone()

		route53domains_deleteTagsForDomainCmd.Flags().String("domain-name", "", "The domain for which you want to delete one or more tags.")
		route53domains_deleteTagsForDomainCmd.Flags().String("tags-to-delete", "", "A list of tag keys to delete.")
		route53domains_deleteTagsForDomainCmd.MarkFlagRequired("domain-name")
		route53domains_deleteTagsForDomainCmd.MarkFlagRequired("tags-to-delete")
	})
	route53domainsCmd.AddCommand(route53domains_deleteTagsForDomainCmd)
}
