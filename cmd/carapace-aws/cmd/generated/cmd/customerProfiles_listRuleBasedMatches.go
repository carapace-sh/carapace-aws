package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_listRuleBasedMatchesCmd = &cobra.Command{
	Use:   "list-rule-based-matches",
	Short: "Returns a set of `MatchIds` that belong to the given domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_listRuleBasedMatchesCmd).Standalone()

	customerProfiles_listRuleBasedMatchesCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_listRuleBasedMatchesCmd.Flags().String("max-results", "", "The maximum number of `MatchIds` returned per page.")
	customerProfiles_listRuleBasedMatchesCmd.Flags().String("next-token", "", "The pagination token from the previous `ListRuleBasedMatches` API call.")
	customerProfiles_listRuleBasedMatchesCmd.MarkFlagRequired("domain-name")
	customerProfilesCmd.AddCommand(customerProfiles_listRuleBasedMatchesCmd)
}
