package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_removeSchemaVersionMetadataCmd = &cobra.Command{
	Use:   "remove-schema-version-metadata",
	Short: "Removes a key value pair from the schema version metadata for the specified schema version ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_removeSchemaVersionMetadataCmd).Standalone()

	glue_removeSchemaVersionMetadataCmd.Flags().String("metadata-key-value", "", "The value of the metadata key.")
	glue_removeSchemaVersionMetadataCmd.Flags().String("schema-id", "", "A wrapper structure that may contain the schema name and Amazon Resource Name (ARN).")
	glue_removeSchemaVersionMetadataCmd.Flags().String("schema-version-id", "", "The unique version ID of the schema version.")
	glue_removeSchemaVersionMetadataCmd.Flags().String("schema-version-number", "", "The version number of the schema.")
	glue_removeSchemaVersionMetadataCmd.MarkFlagRequired("metadata-key-value")
	glueCmd.AddCommand(glue_removeSchemaVersionMetadataCmd)
}
