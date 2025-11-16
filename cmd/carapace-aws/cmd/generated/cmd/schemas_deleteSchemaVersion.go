package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_deleteSchemaVersionCmd = &cobra.Command{
	Use:   "delete-schema-version",
	Short: "Delete the schema version definition",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_deleteSchemaVersionCmd).Standalone()

	schemas_deleteSchemaVersionCmd.Flags().String("registry-name", "", "The name of the registry.")
	schemas_deleteSchemaVersionCmd.Flags().String("schema-name", "", "The name of the schema.")
	schemas_deleteSchemaVersionCmd.Flags().String("schema-version", "", "The version number of the schema")
	schemas_deleteSchemaVersionCmd.MarkFlagRequired("registry-name")
	schemas_deleteSchemaVersionCmd.MarkFlagRequired("schema-name")
	schemas_deleteSchemaVersionCmd.MarkFlagRequired("schema-version")
	schemasCmd.AddCommand(schemas_deleteSchemaVersionCmd)
}
