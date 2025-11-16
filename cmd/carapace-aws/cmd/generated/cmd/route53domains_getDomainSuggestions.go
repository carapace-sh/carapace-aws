package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_getDomainSuggestionsCmd = &cobra.Command{
	Use:   "get-domain-suggestions",
	Short: "The GetDomainSuggestions operation returns a list of suggested domain names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_getDomainSuggestionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_getDomainSuggestionsCmd).Standalone()

		route53domains_getDomainSuggestionsCmd.Flags().String("domain-name", "", "A domain name that you want to use as the basis for a list of possible domain names.")
		route53domains_getDomainSuggestionsCmd.Flags().Bool("no-only-available", false, "If `OnlyAvailable` is `true`, Route 53 returns only domain names that are available.")
		route53domains_getDomainSuggestionsCmd.Flags().Bool("only-available", false, "If `OnlyAvailable` is `true`, Route 53 returns only domain names that are available.")
		route53domains_getDomainSuggestionsCmd.Flags().String("suggestion-count", "", "The number of suggested domain names that you want Route 53 to return.")
		route53domains_getDomainSuggestionsCmd.MarkFlagRequired("domain-name")
		route53domains_getDomainSuggestionsCmd.Flag("no-only-available").Hidden = true
		route53domains_getDomainSuggestionsCmd.MarkFlagRequired("no-only-available")
		route53domains_getDomainSuggestionsCmd.MarkFlagRequired("only-available")
		route53domains_getDomainSuggestionsCmd.MarkFlagRequired("suggestion-count")
	})
	route53domainsCmd.AddCommand(route53domains_getDomainSuggestionsCmd)
}
