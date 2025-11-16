package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_updateTagsForDomainCmd = &cobra.Command{
	Use:   "update-tags-for-domain",
	Short: "This operation adds or updates tags for a specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_updateTagsForDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_updateTagsForDomainCmd).Standalone()

		route53domains_updateTagsForDomainCmd.Flags().String("domain-name", "", "The domain for which you want to add or update tags.")
		route53domains_updateTagsForDomainCmd.Flags().String("tags-to-update", "", "A list of the tag keys and values that you want to add or update.")
		route53domains_updateTagsForDomainCmd.MarkFlagRequired("domain-name")
	})
	route53domainsCmd.AddCommand(route53domains_updateTagsForDomainCmd)
}
