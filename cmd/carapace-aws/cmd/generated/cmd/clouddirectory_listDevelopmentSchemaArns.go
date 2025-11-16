package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listDevelopmentSchemaArnsCmd = &cobra.Command{
	Use:   "list-development-schema-arns",
	Short: "Retrieves each Amazon Resource Name (ARN) of schemas in the development state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listDevelopmentSchemaArnsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_listDevelopmentSchemaArnsCmd).Standalone()

		clouddirectory_listDevelopmentSchemaArnsCmd.Flags().String("max-results", "", "The maximum number of results to retrieve.")
		clouddirectory_listDevelopmentSchemaArnsCmd.Flags().String("next-token", "", "The pagination token.")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_listDevelopmentSchemaArnsCmd)
}
