package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_updateDataSourceCmd = &cobra.Command{
	Use:   "update-data-source",
	Short: "Updates a direct-query data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_updateDataSourceCmd).Standalone()

	opensearch_updateDataSourceCmd.Flags().String("data-source-type", "", "The type of data source.")
	opensearch_updateDataSourceCmd.Flags().String("description", "", "A new description of the data source.")
	opensearch_updateDataSourceCmd.Flags().String("domain-name", "", "The name of the domain.")
	opensearch_updateDataSourceCmd.Flags().String("name", "", "The name of the data source to modify.")
	opensearch_updateDataSourceCmd.Flags().String("status", "", "The status of the data source update.")
	opensearch_updateDataSourceCmd.MarkFlagRequired("data-source-type")
	opensearch_updateDataSourceCmd.MarkFlagRequired("domain-name")
	opensearch_updateDataSourceCmd.MarkFlagRequired("name")
	opensearchCmd.AddCommand(opensearch_updateDataSourceCmd)
}
