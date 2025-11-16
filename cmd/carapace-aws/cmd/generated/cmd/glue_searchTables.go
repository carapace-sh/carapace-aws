package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_searchTablesCmd = &cobra.Command{
	Use:   "search-tables",
	Short: "Searches a set of tables based on properties in the table metadata as well as on the parent database.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_searchTablesCmd).Standalone()

	glue_searchTablesCmd.Flags().String("catalog-id", "", "A unique identifier, consisting of `account_id`.")
	glue_searchTablesCmd.Flags().String("filters", "", "A list of key-value pairs, and a comparator used to filter the search results.")
	glue_searchTablesCmd.Flags().String("include-status-details", "", "Specifies whether to include status details related to a request to create or update an Glue Data Catalog view.")
	glue_searchTablesCmd.Flags().String("max-results", "", "The maximum number of tables to return in a single response.")
	glue_searchTablesCmd.Flags().String("next-token", "", "A continuation token, included if this is a continuation call.")
	glue_searchTablesCmd.Flags().String("resource-share-type", "", "Allows you to specify that you want to search the tables shared with your account.")
	glue_searchTablesCmd.Flags().String("search-text", "", "A string used for a text search.")
	glue_searchTablesCmd.Flags().String("sort-criteria", "", "A list of criteria for sorting the results by a field name, in an ascending or descending order.")
	glueCmd.AddCommand(glue_searchTablesCmd)
}
