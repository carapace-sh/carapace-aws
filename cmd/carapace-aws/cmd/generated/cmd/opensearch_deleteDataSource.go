package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_deleteDataSourceCmd = &cobra.Command{
	Use:   "delete-data-source",
	Short: "Deletes a direct-query data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_deleteDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_deleteDataSourceCmd).Standalone()

		opensearch_deleteDataSourceCmd.Flags().String("domain-name", "", "The name of the domain.")
		opensearch_deleteDataSourceCmd.Flags().String("name", "", "The name of the data source to delete.")
		opensearch_deleteDataSourceCmd.MarkFlagRequired("domain-name")
		opensearch_deleteDataSourceCmd.MarkFlagRequired("name")
	})
	opensearchCmd.AddCommand(opensearch_deleteDataSourceCmd)
}
