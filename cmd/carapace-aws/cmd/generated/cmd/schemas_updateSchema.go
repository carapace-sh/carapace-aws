package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_updateSchemaCmd = &cobra.Command{
	Use:   "update-schema",
	Short: "Updates the schema definition\n\nInactive schemas will be deleted after two years.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_updateSchemaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_updateSchemaCmd).Standalone()

		schemas_updateSchemaCmd.Flags().String("client-token-id", "", "The ID of the client token.")
		schemas_updateSchemaCmd.Flags().String("content", "", "The source of the schema definition.")
		schemas_updateSchemaCmd.Flags().String("description", "", "The description of the schema.")
		schemas_updateSchemaCmd.Flags().String("registry-name", "", "The name of the registry.")
		schemas_updateSchemaCmd.Flags().String("schema-name", "", "The name of the schema.")
		schemas_updateSchemaCmd.Flags().String("type", "", "The schema type for the events schema.")
		schemas_updateSchemaCmd.MarkFlagRequired("registry-name")
		schemas_updateSchemaCmd.MarkFlagRequired("schema-name")
	})
	schemasCmd.AddCommand(schemas_updateSchemaCmd)
}
