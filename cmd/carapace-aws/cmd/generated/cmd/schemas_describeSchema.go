package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_describeSchemaCmd = &cobra.Command{
	Use:   "describe-schema",
	Short: "Retrieve the schema definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_describeSchemaCmd).Standalone()

	schemas_describeSchemaCmd.Flags().String("registry-name", "", "The name of the registry.")
	schemas_describeSchemaCmd.Flags().String("schema-name", "", "The name of the schema.")
	schemas_describeSchemaCmd.Flags().String("schema-version", "", "Specifying this limits the results to only this schema version.")
	schemas_describeSchemaCmd.MarkFlagRequired("registry-name")
	schemas_describeSchemaCmd.MarkFlagRequired("schema-name")
	schemasCmd.AddCommand(schemas_describeSchemaCmd)
}
