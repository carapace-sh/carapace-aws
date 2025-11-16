package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createSchemaCmd = &cobra.Command{
	Use:   "create-schema",
	Short: "Creates an Amazon Personalize schema from the specified schema string.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createSchemaCmd).Standalone()

	personalize_createSchemaCmd.Flags().String("domain", "", "The domain for the schema.")
	personalize_createSchemaCmd.Flags().String("name", "", "The name for the schema.")
	personalize_createSchemaCmd.Flags().String("schema", "", "A schema in Avro JSON format.")
	personalize_createSchemaCmd.MarkFlagRequired("name")
	personalize_createSchemaCmd.MarkFlagRequired("schema")
	personalizeCmd.AddCommand(personalize_createSchemaCmd)
}
