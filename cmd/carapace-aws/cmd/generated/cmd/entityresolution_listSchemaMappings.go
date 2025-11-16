package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_listSchemaMappingsCmd = &cobra.Command{
	Use:   "list-schema-mappings",
	Short: "Returns a list of all the `SchemaMappings` that have been created for an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_listSchemaMappingsCmd).Standalone()

	entityresolution_listSchemaMappingsCmd.Flags().String("max-results", "", "The maximum number of objects returned per page.")
	entityresolution_listSchemaMappingsCmd.Flags().String("next-token", "", "The pagination token from the previous API call.")
	entityresolutionCmd.AddCommand(entityresolution_listSchemaMappingsCmd)
}
