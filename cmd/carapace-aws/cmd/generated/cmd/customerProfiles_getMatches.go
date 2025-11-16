package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getMatchesCmd = &cobra.Command{
	Use:   "get-matches",
	Short: "Before calling this API, use [CreateDomain](https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_CreateDomain.html) or [UpdateDomain](https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UpdateDomain.html) to enable identity resolution: set `Matching` to true.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getMatchesCmd).Standalone()

	customerProfiles_getMatchesCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_getMatchesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	customerProfiles_getMatchesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	customerProfiles_getMatchesCmd.MarkFlagRequired("domain-name")
	customerProfilesCmd.AddCommand(customerProfiles_getMatchesCmd)
}
