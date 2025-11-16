package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listSchemasCmd = &cobra.Command{
	Use:   "list-schemas",
	Short: "Lists the schemas for relations within a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listSchemasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_listSchemasCmd).Standalone()

		cleanrooms_listSchemasCmd.Flags().String("collaboration-identifier", "", "A unique identifier for the collaboration that the schema belongs to.")
		cleanrooms_listSchemasCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
		cleanrooms_listSchemasCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		cleanrooms_listSchemasCmd.Flags().String("schema-type", "", "If present, filter schemas by schema type.")
		cleanrooms_listSchemasCmd.MarkFlagRequired("collaboration-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_listSchemasCmd)
}
