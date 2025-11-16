package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_getDataSourceCmd = &cobra.Command{
	Use:   "get-data-source",
	Short: "Retrieves information about a direct query data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_getDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_getDataSourceCmd).Standalone()

		opensearch_getDataSourceCmd.Flags().String("domain-name", "", "The name of the domain.")
		opensearch_getDataSourceCmd.Flags().String("name", "", "The name of the data source to get information about.")
		opensearch_getDataSourceCmd.MarkFlagRequired("domain-name")
		opensearch_getDataSourceCmd.MarkFlagRequired("name")
	})
	opensearchCmd.AddCommand(opensearch_getDataSourceCmd)
}
