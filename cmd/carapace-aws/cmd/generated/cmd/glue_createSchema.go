package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createSchemaCmd = &cobra.Command{
	Use:   "create-schema",
	Short: "Creates a new schema set and registers the schema definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createSchemaCmd).Standalone()

	glue_createSchemaCmd.Flags().String("compatibility", "", "The compatibility mode of the schema.")
	glue_createSchemaCmd.Flags().String("data-format", "", "The data format of the schema definition.")
	glue_createSchemaCmd.Flags().String("description", "", "An optional description of the schema.")
	glue_createSchemaCmd.Flags().String("registry-id", "", "This is a wrapper shape to contain the registry identity fields.")
	glue_createSchemaCmd.Flags().String("schema-definition", "", "The schema definition using the `DataFormat` setting for `SchemaName`.")
	glue_createSchemaCmd.Flags().String("schema-name", "", "Name of the schema to be created of max length of 255, and may only contain letters, numbers, hyphen, underscore, dollar sign, or hash mark.")
	glue_createSchemaCmd.Flags().String("tags", "", "Amazon Web Services tags that contain a key value pair and may be searched by console, command line, or API.")
	glue_createSchemaCmd.MarkFlagRequired("data-format")
	glue_createSchemaCmd.MarkFlagRequired("schema-name")
	glueCmd.AddCommand(glue_createSchemaCmd)
}
