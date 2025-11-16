package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_putSchemaVersionMetadataCmd = &cobra.Command{
	Use:   "put-schema-version-metadata",
	Short: "Puts the metadata key value pair for a specified schema version ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_putSchemaVersionMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_putSchemaVersionMetadataCmd).Standalone()

		glue_putSchemaVersionMetadataCmd.Flags().String("metadata-key-value", "", "The metadata key's corresponding value.")
		glue_putSchemaVersionMetadataCmd.Flags().String("schema-id", "", "The unique ID for the schema.")
		glue_putSchemaVersionMetadataCmd.Flags().String("schema-version-id", "", "The unique version ID of the schema version.")
		glue_putSchemaVersionMetadataCmd.Flags().String("schema-version-number", "", "The version number of the schema.")
		glue_putSchemaVersionMetadataCmd.MarkFlagRequired("metadata-key-value")
	})
	glueCmd.AddCommand(glue_putSchemaVersionMetadataCmd)
}
