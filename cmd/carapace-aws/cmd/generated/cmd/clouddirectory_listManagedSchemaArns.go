package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clouddirectory_listManagedSchemaArnsCmd = &cobra.Command{
	Use:   "list-managed-schema-arns",
	Short: "Lists the major version families of each managed schema.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clouddirectory_listManagedSchemaArnsCmd).Standalone()

	clouddirectory_listManagedSchemaArnsCmd.Flags().String("max-results", "", "The maximum number of results to retrieve.")
	clouddirectory_listManagedSchemaArnsCmd.Flags().String("next-token", "", "The pagination token.")
	clouddirectory_listManagedSchemaArnsCmd.Flags().String("schema-arn", "", "The response for ListManagedSchemaArns.")
	clouddirectoryCmd.AddCommand(clouddirectory_listManagedSchemaArnsCmd)
}
