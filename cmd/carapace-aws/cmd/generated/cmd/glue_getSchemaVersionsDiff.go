package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getSchemaVersionsDiffCmd = &cobra.Command{
	Use:   "get-schema-versions-diff",
	Short: "Fetches the schema version difference in the specified difference type between two stored schema versions in the Schema Registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getSchemaVersionsDiffCmd).Standalone()

	glue_getSchemaVersionsDiffCmd.Flags().String("first-schema-version-number", "", "The first of the two schema versions to be compared.")
	glue_getSchemaVersionsDiffCmd.Flags().String("schema-diff-type", "", "Refers to `SYNTAX_DIFF`, which is the currently supported diff type.")
	glue_getSchemaVersionsDiffCmd.Flags().String("schema-id", "", "This is a wrapper structure to contain schema identity fields.")
	glue_getSchemaVersionsDiffCmd.Flags().String("second-schema-version-number", "", "The second of the two schema versions to be compared.")
	glue_getSchemaVersionsDiffCmd.MarkFlagRequired("first-schema-version-number")
	glue_getSchemaVersionsDiffCmd.MarkFlagRequired("schema-diff-type")
	glue_getSchemaVersionsDiffCmd.MarkFlagRequired("schema-id")
	glue_getSchemaVersionsDiffCmd.MarkFlagRequired("second-schema-version-number")
	glueCmd.AddCommand(glue_getSchemaVersionsDiffCmd)
}
