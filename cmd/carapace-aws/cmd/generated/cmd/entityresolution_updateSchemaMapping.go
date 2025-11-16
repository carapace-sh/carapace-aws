package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_updateSchemaMappingCmd = &cobra.Command{
	Use:   "update-schema-mapping",
	Short: "Updates a schema mapping.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_updateSchemaMappingCmd).Standalone()

	entityresolution_updateSchemaMappingCmd.Flags().String("description", "", "A description of the schema.")
	entityresolution_updateSchemaMappingCmd.Flags().String("mapped-input-fields", "", "A list of `MappedInputFields`.")
	entityresolution_updateSchemaMappingCmd.Flags().String("schema-name", "", "The name of the schema.")
	entityresolution_updateSchemaMappingCmd.MarkFlagRequired("mapped-input-fields")
	entityresolution_updateSchemaMappingCmd.MarkFlagRequired("schema-name")
	entityresolutionCmd.AddCommand(entityresolution_updateSchemaMappingCmd)
}
