package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_listDataSourcesCmd = &cobra.Command{
	Use:   "list-data-sources",
	Short: "Lists the Amazon Q Business data source connectors that you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_listDataSourcesCmd).Standalone()

	qbusiness_listDataSourcesCmd.Flags().String("application-id", "", "The identifier of the Amazon Q Business application linked to the data source connectors.")
	qbusiness_listDataSourcesCmd.Flags().String("index-id", "", "The identifier of the index used with one or more data source connectors.")
	qbusiness_listDataSourcesCmd.Flags().String("max-results", "", "The maximum number of data source connectors to return.")
	qbusiness_listDataSourcesCmd.Flags().String("next-token", "", "If the `maxResults` response was incomplete because there is more data to retrieve, Amazon Q Business returns a pagination token in the response.")
	qbusiness_listDataSourcesCmd.MarkFlagRequired("application-id")
	qbusiness_listDataSourcesCmd.MarkFlagRequired("index-id")
	qbusinessCmd.AddCommand(qbusiness_listDataSourcesCmd)
}
