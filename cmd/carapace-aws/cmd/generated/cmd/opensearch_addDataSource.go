package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_addDataSourceCmd = &cobra.Command{
	Use:   "add-data-source",
	Short: "Creates a new direct-query data source to the specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_addDataSourceCmd).Standalone()

	opensearch_addDataSourceCmd.Flags().String("data-source-type", "", "The type of data source.")
	opensearch_addDataSourceCmd.Flags().String("description", "", "A description of the data source.")
	opensearch_addDataSourceCmd.Flags().String("domain-name", "", "The name of the domain to add the data source to.")
	opensearch_addDataSourceCmd.Flags().String("name", "", "A name for the data source.")
	opensearch_addDataSourceCmd.MarkFlagRequired("data-source-type")
	opensearch_addDataSourceCmd.MarkFlagRequired("domain-name")
	opensearch_addDataSourceCmd.MarkFlagRequired("name")
	opensearchCmd.AddCommand(opensearch_addDataSourceCmd)
}
