package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listDataSourceRunActivitiesCmd = &cobra.Command{
	Use:   "list-data-source-run-activities",
	Short: "Lists data source run activities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listDataSourceRunActivitiesCmd).Standalone()

	datazone_listDataSourceRunActivitiesCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which to list data source run activities.")
	datazone_listDataSourceRunActivitiesCmd.Flags().String("identifier", "", "The identifier of the data source run.")
	datazone_listDataSourceRunActivitiesCmd.Flags().String("max-results", "", "The maximum number of activities to return in a single call to `ListDataSourceRunActivities`.")
	datazone_listDataSourceRunActivitiesCmd.Flags().String("next-token", "", "When the number of activities is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of activities, the response includes a pagination token named `NextToken`.")
	datazone_listDataSourceRunActivitiesCmd.Flags().String("status", "", "The status of the data source run.")
	datazone_listDataSourceRunActivitiesCmd.MarkFlagRequired("domain-identifier")
	datazone_listDataSourceRunActivitiesCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_listDataSourceRunActivitiesCmd)
}
