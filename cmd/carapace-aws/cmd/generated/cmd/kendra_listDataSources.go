package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_listDataSourcesCmd = &cobra.Command{
	Use:   "list-data-sources",
	Short: "Lists the data source connectors that you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_listDataSourcesCmd).Standalone()

	kendra_listDataSourcesCmd.Flags().String("index-id", "", "The identifier of the index used with one or more data source connectors.")
	kendra_listDataSourcesCmd.Flags().String("max-results", "", "The maximum number of data source connectors to return.")
	kendra_listDataSourcesCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there is more data to retrieve), Amazon Kendra returns a pagination token in the response.")
	kendra_listDataSourcesCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_listDataSourcesCmd)
}
