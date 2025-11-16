package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_createSchemaCmd = &cobra.Command{
	Use:   "create-schema",
	Short: "Creates a schema definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_createSchemaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_createSchemaCmd).Standalone()

		schemas_createSchemaCmd.Flags().String("content", "", "The source of the schema definition.")
		schemas_createSchemaCmd.Flags().String("description", "", "A description of the schema.")
		schemas_createSchemaCmd.Flags().String("registry-name", "", "The name of the registry.")
		schemas_createSchemaCmd.Flags().String("schema-name", "", "The name of the schema.")
		schemas_createSchemaCmd.Flags().String("tags", "", "Tags associated with the schema.")
		schemas_createSchemaCmd.Flags().String("type", "", "The type of schema.")
		schemas_createSchemaCmd.MarkFlagRequired("content")
		schemas_createSchemaCmd.MarkFlagRequired("registry-name")
		schemas_createSchemaCmd.MarkFlagRequired("schema-name")
		schemas_createSchemaCmd.MarkFlagRequired("type")
	})
	schemasCmd.AddCommand(schemas_createSchemaCmd)
}
