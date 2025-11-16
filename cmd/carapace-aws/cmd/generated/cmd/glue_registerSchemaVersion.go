package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_registerSchemaVersionCmd = &cobra.Command{
	Use:   "register-schema-version",
	Short: "Adds a new version to the existing schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_registerSchemaVersionCmd).Standalone()

	glue_registerSchemaVersionCmd.Flags().String("schema-definition", "", "The schema definition using the `DataFormat` setting for the `SchemaName`.")
	glue_registerSchemaVersionCmd.Flags().String("schema-id", "", "This is a wrapper structure to contain schema identity fields.")
	glue_registerSchemaVersionCmd.MarkFlagRequired("schema-definition")
	glue_registerSchemaVersionCmd.MarkFlagRequired("schema-id")
	glueCmd.AddCommand(glue_registerSchemaVersionCmd)
}
