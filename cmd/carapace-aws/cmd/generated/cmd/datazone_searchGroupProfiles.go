package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_searchGroupProfilesCmd = &cobra.Command{
	Use:   "search-group-profiles",
	Short: "Searches group profiles in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_searchGroupProfilesCmd).Standalone()

	datazone_searchGroupProfilesCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which you want to search group profiles.")
	datazone_searchGroupProfilesCmd.Flags().String("group-type", "", "The group type for which to search.")
	datazone_searchGroupProfilesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call to `SearchGroupProfiles`.")
	datazone_searchGroupProfilesCmd.Flags().String("next-token", "", "When the number of results is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of results, the response includes a pagination token named `NextToken`.")
	datazone_searchGroupProfilesCmd.Flags().String("search-text", "", "Specifies the text for which to search.")
	datazone_searchGroupProfilesCmd.MarkFlagRequired("domain-identifier")
	datazone_searchGroupProfilesCmd.MarkFlagRequired("group-type")
	datazoneCmd.AddCommand(datazone_searchGroupProfilesCmd)
}
