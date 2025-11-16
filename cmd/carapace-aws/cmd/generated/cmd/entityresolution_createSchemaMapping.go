package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_createSchemaMappingCmd = &cobra.Command{
	Use:   "create-schema-mapping",
	Short: "Creates a schema mapping, which defines the schema of the input customer records table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_createSchemaMappingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_createSchemaMappingCmd).Standalone()

		entityresolution_createSchemaMappingCmd.Flags().String("description", "", "A description of the schema.")
		entityresolution_createSchemaMappingCmd.Flags().String("mapped-input-fields", "", "A list of `MappedInputFields`.")
		entityresolution_createSchemaMappingCmd.Flags().String("schema-name", "", "The name of the schema.")
		entityresolution_createSchemaMappingCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
		entityresolution_createSchemaMappingCmd.MarkFlagRequired("mapped-input-fields")
		entityresolution_createSchemaMappingCmd.MarkFlagRequired("schema-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_createSchemaMappingCmd)
}
