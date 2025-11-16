package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listDataSourcesCmd = &cobra.Command{
	Use:   "list-data-sources",
	Short: "Lists data sources in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listDataSourcesCmd).Standalone()

	datazone_listDataSourcesCmd.Flags().String("connection-identifier", "", "The ID of the connection.")
	datazone_listDataSourcesCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which to list the data sources.")
	datazone_listDataSourcesCmd.Flags().String("environment-identifier", "", "The identifier of the environment in which to list the data sources.")
	datazone_listDataSourcesCmd.Flags().String("max-results", "", "The maximum number of data sources to return in a single call to `ListDataSources`.")
	datazone_listDataSourcesCmd.Flags().String("name", "", "The name of the data source.")
	datazone_listDataSourcesCmd.Flags().String("next-token", "", "When the number of data sources is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of data sources, the response includes a pagination token named `NextToken`.")
	datazone_listDataSourcesCmd.Flags().String("project-identifier", "", "The identifier of the project in which to list data sources.")
	datazone_listDataSourcesCmd.Flags().String("status", "", "The status of the data source.")
	datazone_listDataSourcesCmd.Flags().String("type", "", "The type of the data source.")
	datazone_listDataSourcesCmd.MarkFlagRequired("domain-identifier")
	datazone_listDataSourcesCmd.MarkFlagRequired("project-identifier")
	datazoneCmd.AddCommand(datazone_listDataSourcesCmd)
}
