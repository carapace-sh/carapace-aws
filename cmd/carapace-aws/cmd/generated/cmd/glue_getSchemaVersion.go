package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getSchemaVersionCmd = &cobra.Command{
	Use:   "get-schema-version",
	Short: "Get the specified schema by its unique ID assigned when a version of the schema is created or registered.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getSchemaVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_getSchemaVersionCmd).Standalone()

		glue_getSchemaVersionCmd.Flags().String("schema-id", "", "This is a wrapper structure to contain schema identity fields.")
		glue_getSchemaVersionCmd.Flags().String("schema-version-id", "", "The `SchemaVersionId` of the schema version.")
		glue_getSchemaVersionCmd.Flags().String("schema-version-number", "", "The version number of the schema.")
	})
	glueCmd.AddCommand(glue_getSchemaVersionCmd)
}
