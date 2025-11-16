package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_deleteSchemaMappingCmd = &cobra.Command{
	Use:   "delete-schema-mapping",
	Short: "Deletes the `SchemaMapping` with a given name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_deleteSchemaMappingCmd).Standalone()

	entityresolution_deleteSchemaMappingCmd.Flags().String("schema-name", "", "The name of the schema to delete.")
	entityresolution_deleteSchemaMappingCmd.MarkFlagRequired("schema-name")
	entityresolutionCmd.AddCommand(entityresolution_deleteSchemaMappingCmd)
}
