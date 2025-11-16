package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listPublishedSchemaArnsCmd = &cobra.Command{
	Use:   "list-published-schema-arns",
	Short: "Lists the major version families of each published schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listPublishedSchemaArnsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(clouddirectory_listPublishedSchemaArnsCmd).Standalone()

		clouddirectory_listPublishedSchemaArnsCmd.Flags().String("max-results", "", "The maximum number of results to retrieve.")
		clouddirectory_listPublishedSchemaArnsCmd.Flags().String("next-token", "", "The pagination token.")
		clouddirectory_listPublishedSchemaArnsCmd.Flags().String("schema-arn", "", "The response for `ListPublishedSchemaArns` when this parameter is used will list all minor version ARNs for a major version.")
	})
	clouddirectoryCmd.AddCommand(clouddirectory_listPublishedSchemaArnsCmd)
}
