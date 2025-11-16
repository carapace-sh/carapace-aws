package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_deleteSchemaCmd = &cobra.Command{
	Use:   "delete-schema",
	Short: "Delete a schema definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_deleteSchemaCmd).Standalone()

	schemas_deleteSchemaCmd.Flags().String("registry-name", "", "The name of the registry.")
	schemas_deleteSchemaCmd.Flags().String("schema-name", "", "The name of the schema.")
	schemas_deleteSchemaCmd.MarkFlagRequired("registry-name")
	schemas_deleteSchemaCmd.MarkFlagRequired("schema-name")
	schemasCmd.AddCommand(schemas_deleteSchemaCmd)
}
