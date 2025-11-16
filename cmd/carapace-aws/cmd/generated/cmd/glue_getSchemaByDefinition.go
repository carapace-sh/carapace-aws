package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getSchemaByDefinitionCmd = &cobra.Command{
	Use:   "get-schema-by-definition",
	Short: "Retrieves a schema by the `SchemaDefinition`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getSchemaByDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getSchemaByDefinitionCmd).Standalone()

		glue_getSchemaByDefinitionCmd.Flags().String("schema-definition", "", "The definition of the schema for which schema details are required.")
		glue_getSchemaByDefinitionCmd.Flags().String("schema-id", "", "This is a wrapper structure to contain schema identity fields.")
		glue_getSchemaByDefinitionCmd.MarkFlagRequired("schema-definition")
		glue_getSchemaByDefinitionCmd.MarkFlagRequired("schema-id")
	})
	glueCmd.AddCommand(glue_getSchemaByDefinitionCmd)
}
