package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_listSchemasCmd = &cobra.Command{
	Use:   "list-schemas",
	Short: "List the schemas.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_listSchemasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_listSchemasCmd).Standalone()

		schemas_listSchemasCmd.Flags().String("limit", "", "")
		schemas_listSchemasCmd.Flags().String("next-token", "", "The token that specifies the next page of results to return.")
		schemas_listSchemasCmd.Flags().String("registry-name", "", "The name of the registry.")
		schemas_listSchemasCmd.Flags().String("schema-name-prefix", "", "Specifying this limits the results to only those schema names that start with the specified prefix.")
		schemas_listSchemasCmd.MarkFlagRequired("registry-name")
	})
	schemasCmd.AddCommand(schemas_listSchemasCmd)
}
