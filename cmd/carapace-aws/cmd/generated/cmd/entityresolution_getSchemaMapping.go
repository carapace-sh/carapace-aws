package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_getSchemaMappingCmd = &cobra.Command{
	Use:   "get-schema-mapping",
	Short: "Returns the SchemaMapping of a given name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_getSchemaMappingCmd).Standalone()

	entityresolution_getSchemaMappingCmd.Flags().String("schema-name", "", "The name of the schema to be retrieved.")
	entityresolution_getSchemaMappingCmd.MarkFlagRequired("schema-name")
	entityresolutionCmd.AddCommand(entityresolution_getSchemaMappingCmd)
}
