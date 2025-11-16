package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_searchSchemasCmd = &cobra.Command{
	Use:   "search-schemas",
	Short: "Search the schemas",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_searchSchemasCmd).Standalone()

	schemas_searchSchemasCmd.Flags().String("keywords", "", "Specifying this limits the results to only schemas that include the provided keywords.")
	schemas_searchSchemasCmd.Flags().String("limit", "", "")
	schemas_searchSchemasCmd.Flags().String("next-token", "", "The token that specifies the next page of results to return.")
	schemas_searchSchemasCmd.Flags().String("registry-name", "", "The name of the registry.")
	schemas_searchSchemasCmd.MarkFlagRequired("keywords")
	schemas_searchSchemasCmd.MarkFlagRequired("registry-name")
	schemasCmd.AddCommand(schemas_searchSchemasCmd)
}
