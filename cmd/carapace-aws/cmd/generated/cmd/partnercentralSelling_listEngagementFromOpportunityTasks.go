package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_listEngagementFromOpportunityTasksCmd = &cobra.Command{
	Use:   "list-engagement-from-opportunity-tasks",
	Short: "Lists all in-progress, completed, or failed `EngagementFromOpportunity` tasks that were initiated by the caller's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_listEngagementFromOpportunityTasksCmd).Standalone()

	partnercentralSelling_listEngagementFromOpportunityTasksCmd.Flags().String("catalog", "", "Specifies the catalog related to the request.")
	partnercentralSelling_listEngagementFromOpportunityTasksCmd.Flags().String("engagement-identifier", "", "Filters tasks by the identifiers of the engagements they created or are associated with.")
	partnercentralSelling_listEngagementFromOpportunityTasksCmd.Flags().String("max-results", "", "Specifies the maximum number of results to return in a single page of the response.Use this parameter to control the number of items returned in each request, which can be useful for performance tuning and managing large result sets.")
	partnercentralSelling_listEngagementFromOpportunityTasksCmd.Flags().String("next-token", "", "The token for requesting the next page of results.")
	partnercentralSelling_listEngagementFromOpportunityTasksCmd.Flags().String("opportunity-identifier", "", "The identifier of the original opportunity associated with this task.")
	partnercentralSelling_listEngagementFromOpportunityTasksCmd.Flags().String("sort", "", "Specifies the sorting criteria for the returned results.")
	partnercentralSelling_listEngagementFromOpportunityTasksCmd.Flags().String("task-identifier", "", "Filters tasks by their unique identifiers.")
	partnercentralSelling_listEngagementFromOpportunityTasksCmd.Flags().String("task-status", "", "Filters the tasks based on their current status.")
	partnercentralSelling_listEngagementFromOpportunityTasksCmd.MarkFlagRequired("catalog")
	partnercentralSellingCmd.AddCommand(partnercentralSelling_listEngagementFromOpportunityTasksCmd)
}
