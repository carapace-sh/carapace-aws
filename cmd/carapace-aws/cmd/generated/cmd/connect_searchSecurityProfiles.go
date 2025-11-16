package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchSecurityProfilesCmd = &cobra.Command{
	Use:   "search-security-profiles",
	Short: "Searches security profiles in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchSecurityProfilesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_searchSecurityProfilesCmd).Standalone()

		connect_searchSecurityProfilesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_searchSecurityProfilesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_searchSecurityProfilesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_searchSecurityProfilesCmd.Flags().String("search-criteria", "", "The search criteria to be used to return security profiles.")
		connect_searchSecurityProfilesCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
		connect_searchSecurityProfilesCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_searchSecurityProfilesCmd)
}
