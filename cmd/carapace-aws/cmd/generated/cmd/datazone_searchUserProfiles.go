package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_searchUserProfilesCmd = &cobra.Command{
	Use:   "search-user-profiles",
	Short: "Searches user profiles in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_searchUserProfilesCmd).Standalone()

	datazone_searchUserProfilesCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which you want to search user profiles.")
	datazone_searchUserProfilesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call to `SearchUserProfiles`.")
	datazone_searchUserProfilesCmd.Flags().String("next-token", "", "When the number of results is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of results, the response includes a pagination token named `NextToken`.")
	datazone_searchUserProfilesCmd.Flags().String("search-text", "", "Specifies the text for which to search.")
	datazone_searchUserProfilesCmd.Flags().String("user-type", "", "Specifies the user type for the `SearchUserProfiles` action.")
	datazone_searchUserProfilesCmd.MarkFlagRequired("domain-identifier")
	datazone_searchUserProfilesCmd.MarkFlagRequired("user-type")
	datazoneCmd.AddCommand(datazone_searchUserProfilesCmd)
}
