package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_updateSchemaCmd = &cobra.Command{
	Use:   "update-schema",
	Short: "Updates the description, compatibility setting, or version checkpoint for a schema set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_updateSchemaCmd).Standalone()

	glue_updateSchemaCmd.Flags().String("compatibility", "", "The new compatibility setting for the schema.")
	glue_updateSchemaCmd.Flags().String("description", "", "The new description for the schema.")
	glue_updateSchemaCmd.Flags().String("schema-id", "", "This is a wrapper structure to contain schema identity fields.")
	glue_updateSchemaCmd.Flags().String("schema-version-number", "", "Version number required for check pointing.")
	glue_updateSchemaCmd.MarkFlagRequired("schema-id")
	glueCmd.AddCommand(glue_updateSchemaCmd)
}
