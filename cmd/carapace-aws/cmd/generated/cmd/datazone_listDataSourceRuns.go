package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listDataSourceRunsCmd = &cobra.Command{
	Use:   "list-data-source-runs",
	Short: "Lists data source runs in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listDataSourceRunsCmd).Standalone()

	datazone_listDataSourceRunsCmd.Flags().String("data-source-identifier", "", "The identifier of the data source.")
	datazone_listDataSourceRunsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which to invoke the `ListDataSourceRuns` action.")
	datazone_listDataSourceRunsCmd.Flags().String("max-results", "", "The maximum number of runs to return in a single call to `ListDataSourceRuns`.")
	datazone_listDataSourceRunsCmd.Flags().String("next-token", "", "When the number of runs is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of runs, the response includes a pagination token named `NextToken`.")
	datazone_listDataSourceRunsCmd.Flags().String("status", "", "The status of the data source.")
	datazone_listDataSourceRunsCmd.MarkFlagRequired("data-source-identifier")
	datazone_listDataSourceRunsCmd.MarkFlagRequired("domain-identifier")
	datazoneCmd.AddCommand(datazone_listDataSourceRunsCmd)
}
