package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listSchemasCmd = &cobra.Command{
	Use:   "list-schemas",
	Short: "Returns a list of schemas with minimal details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listSchemasCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_listSchemasCmd).Standalone()

		glue_listSchemasCmd.Flags().String("max-results", "", "Maximum number of results required per page.")
		glue_listSchemasCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
		glue_listSchemasCmd.Flags().String("registry-id", "", "A wrapper structure that may contain the registry name and Amazon Resource Name (ARN).")
	})
	glueCmd.AddCommand(glue_listSchemasCmd)
}
