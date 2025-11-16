package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listLineageNodeHistoryCmd = &cobra.Command{
	Use:   "list-lineage-node-history",
	Short: "Lists the history of the specified data lineage node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listLineageNodeHistoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_listLineageNodeHistoryCmd).Standalone()

		datazone_listLineageNodeHistoryCmd.Flags().String("direction", "", "The direction of the data lineage node refers to the lineage node having neighbors in that direction.")
		datazone_listLineageNodeHistoryCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to list the history of the specified data lineage node.")
		datazone_listLineageNodeHistoryCmd.Flags().String("event-timestamp-gte", "", "Specifies whether the action is to return data lineage node history from the time after the event timestamp.")
		datazone_listLineageNodeHistoryCmd.Flags().String("event-timestamp-lte", "", "Specifies whether the action is to return data lineage node history from the time prior of the event timestamp.")
		datazone_listLineageNodeHistoryCmd.Flags().String("identifier", "", "The ID of the data lineage node whose history you want to list.")
		datazone_listLineageNodeHistoryCmd.Flags().String("max-results", "", "The maximum number of history items to return in a single call to ListLineageNodeHistory.")
		datazone_listLineageNodeHistoryCmd.Flags().String("next-token", "", "When the number of history items is greater than the default value for the MaxResults parameter, or if you explicitly specify a value for MaxResults that is less than the number of items, the response includes a pagination token named NextToken.")
		datazone_listLineageNodeHistoryCmd.Flags().String("sort-order", "", "The order by which you want data lineage node history to be sorted.")
		datazone_listLineageNodeHistoryCmd.MarkFlagRequired("domain-identifier")
		datazone_listLineageNodeHistoryCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_listLineageNodeHistoryCmd)
}
