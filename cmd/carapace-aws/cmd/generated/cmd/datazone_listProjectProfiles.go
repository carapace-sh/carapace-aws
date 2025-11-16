package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listProjectProfilesCmd = &cobra.Command{
	Use:   "list-project-profiles",
	Short: "Lists project profiles.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listProjectProfilesCmd).Standalone()

	datazone_listProjectProfilesCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to list project profiles.")
	datazone_listProjectProfilesCmd.Flags().String("max-results", "", "The maximum number of project profiles to return in a single call to ListProjectProfiles.")
	datazone_listProjectProfilesCmd.Flags().String("name", "", "The name of a project profile.")
	datazone_listProjectProfilesCmd.Flags().String("next-token", "", "When the number of project profiles is greater than the default value for the MaxResults parameter, or if you explicitly specify a value for MaxResults that is less than the number of project profiles, the response includes a pagination token named NextToken.")
	datazone_listProjectProfilesCmd.Flags().String("sort-by", "", "Specifies by what to sort project profiles.")
	datazone_listProjectProfilesCmd.Flags().String("sort-order", "", "Specifies the sort order of the project profiles.")
	datazone_listProjectProfilesCmd.MarkFlagRequired("domain-identifier")
	datazoneCmd.AddCommand(datazone_listProjectProfilesCmd)
}
