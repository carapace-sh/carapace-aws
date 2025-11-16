package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_listSchemaVersionsCmd = &cobra.Command{
	Use:   "list-schema-versions",
	Short: "Provides a list of the schema versions and related information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_listSchemaVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_listSchemaVersionsCmd).Standalone()

		schemas_listSchemaVersionsCmd.Flags().String("limit", "", "")
		schemas_listSchemaVersionsCmd.Flags().String("next-token", "", "The token that specifies the next page of results to return.")
		schemas_listSchemaVersionsCmd.Flags().String("registry-name", "", "The name of the registry.")
		schemas_listSchemaVersionsCmd.Flags().String("schema-name", "", "The name of the schema.")
		schemas_listSchemaVersionsCmd.MarkFlagRequired("registry-name")
		schemas_listSchemaVersionsCmd.MarkFlagRequired("schema-name")
	})
	schemasCmd.AddCommand(schemas_listSchemaVersionsCmd)
}
