package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var customerProfiles_getSimilarProfilesCmd = &cobra.Command{
	Use:   "get-similar-profiles",
	Short: "Returns a set of profiles that belong to the same matching group using the `matchId` or `profileId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customerProfiles_getSimilarProfilesCmd).Standalone()

	customerProfiles_getSimilarProfilesCmd.Flags().String("domain-name", "", "The unique name of the domain.")
	customerProfiles_getSimilarProfilesCmd.Flags().String("match-type", "", "Specify the type of matching to get similar profiles for.")
	customerProfiles_getSimilarProfilesCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
	customerProfiles_getSimilarProfilesCmd.Flags().String("next-token", "", "The pagination token from the previous `GetSimilarProfiles` API call.")
	customerProfiles_getSimilarProfilesCmd.Flags().String("search-key", "", "The string indicating the search key to be used.")
	customerProfiles_getSimilarProfilesCmd.Flags().String("search-value", "", "The string based on `SearchKey` to be searched for similar profiles.")
	customerProfiles_getSimilarProfilesCmd.MarkFlagRequired("domain-name")
	customerProfiles_getSimilarProfilesCmd.MarkFlagRequired("match-type")
	customerProfiles_getSimilarProfilesCmd.MarkFlagRequired("search-key")
	customerProfiles_getSimilarProfilesCmd.MarkFlagRequired("search-value")
	customerProfilesCmd.AddCommand(customerProfiles_getSimilarProfilesCmd)
}
