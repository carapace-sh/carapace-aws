package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_checkSchemaVersionValidityCmd = &cobra.Command{
	Use:   "check-schema-version-validity",
	Short: "Validates the supplied schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_checkSchemaVersionValidityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_checkSchemaVersionValidityCmd).Standalone()

		glue_checkSchemaVersionValidityCmd.Flags().String("data-format", "", "The data format of the schema definition.")
		glue_checkSchemaVersionValidityCmd.Flags().String("schema-definition", "", "The definition of the schema that has to be validated.")
		glue_checkSchemaVersionValidityCmd.MarkFlagRequired("data-format")
		glue_checkSchemaVersionValidityCmd.MarkFlagRequired("schema-definition")
	})
	glueCmd.AddCommand(glue_checkSchemaVersionValidityCmd)
}
