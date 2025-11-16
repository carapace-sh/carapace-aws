package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workdocs_searchResourcesCmd = &cobra.Command{
	Use:   "search-resources",
	Short: "Searches metadata and the content of folders, documents, document versions, and comments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workdocs_searchResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workdocs_searchResourcesCmd).Standalone()

		workdocs_searchResourcesCmd.Flags().String("additional-response-fields", "", "A list of attributes to include in the response.")
		workdocs_searchResourcesCmd.Flags().String("authentication-token", "", "Amazon WorkDocs authentication token.")
		workdocs_searchResourcesCmd.Flags().String("filters", "", "Filters results based on entity metadata.")
		workdocs_searchResourcesCmd.Flags().String("limit", "", "Max results count per page.")
		workdocs_searchResourcesCmd.Flags().String("marker", "", "The marker for the next set of results.")
		workdocs_searchResourcesCmd.Flags().String("order-by", "", "Order by results in one or more categories.")
		workdocs_searchResourcesCmd.Flags().String("organization-id", "", "Filters based on the resource owner OrgId.")
		workdocs_searchResourcesCmd.Flags().String("query-scopes", "", "Filter based on the text field type.")
		workdocs_searchResourcesCmd.Flags().String("query-text", "", "The String to search for.")
	})
	workdocsCmd.AddCommand(workdocs_searchResourcesCmd)
}
