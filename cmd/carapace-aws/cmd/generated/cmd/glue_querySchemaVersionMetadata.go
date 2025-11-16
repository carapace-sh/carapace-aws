package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_querySchemaVersionMetadataCmd = &cobra.Command{
	Use:   "query-schema-version-metadata",
	Short: "Queries for the schema version metadata information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_querySchemaVersionMetadataCmd).Standalone()

	glue_querySchemaVersionMetadataCmd.Flags().String("max-results", "", "Maximum number of results required per page.")
	glue_querySchemaVersionMetadataCmd.Flags().String("metadata-list", "", "Search key-value pairs for metadata, if they are not provided all the metadata information will be fetched.")
	glue_querySchemaVersionMetadataCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
	glue_querySchemaVersionMetadataCmd.Flags().String("schema-id", "", "A wrapper structure that may contain the schema name and Amazon Resource Name (ARN).")
	glue_querySchemaVersionMetadataCmd.Flags().String("schema-version-id", "", "The unique version ID of the schema version.")
	glue_querySchemaVersionMetadataCmd.Flags().String("schema-version-number", "", "The version number of the schema.")
	glueCmd.AddCommand(glue_querySchemaVersionMetadataCmd)
}
