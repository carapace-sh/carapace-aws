package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteSchemaVersionsCmd = &cobra.Command{
	Use:   "delete-schema-versions",
	Short: "Remove versions from the specified schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteSchemaVersionsCmd).Standalone()

	glue_deleteSchemaVersionsCmd.Flags().String("schema-id", "", "This is a wrapper structure that may contain the schema name and Amazon Resource Name (ARN).")
	glue_deleteSchemaVersionsCmd.Flags().String("versions", "", "A version range may be supplied which may be of the format:")
	glue_deleteSchemaVersionsCmd.MarkFlagRequired("schema-id")
	glue_deleteSchemaVersionsCmd.MarkFlagRequired("versions")
	glueCmd.AddCommand(glue_deleteSchemaVersionsCmd)
}
