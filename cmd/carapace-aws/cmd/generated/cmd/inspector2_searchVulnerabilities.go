package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_searchVulnerabilitiesCmd = &cobra.Command{
	Use:   "search-vulnerabilities",
	Short: "Lists Amazon Inspector coverage details for a specific vulnerability.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_searchVulnerabilitiesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_searchVulnerabilitiesCmd).Standalone()

		inspector2_searchVulnerabilitiesCmd.Flags().String("filter-criteria", "", "The criteria used to filter the results of a vulnerability search.")
		inspector2_searchVulnerabilitiesCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
		inspector2_searchVulnerabilitiesCmd.MarkFlagRequired("filter-criteria")
	})
	inspector2Cmd.AddCommand(inspector2_searchVulnerabilitiesCmd)
}
