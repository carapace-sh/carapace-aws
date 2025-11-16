package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd = &cobra.Command{
	Use:   "list-engagement-by-accepting-invitation-tasks",
	Short: "Lists all in-progress, completed, or failed StartEngagementByAcceptingInvitationTask tasks that were initiated by the caller's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd).Standalone()

		partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd.Flags().String("catalog", "", "Specifies the catalog related to the request.")
		partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd.Flags().String("engagement-invitation-identifier", "", "Filters tasks by the identifiers of the engagement invitations they are processing.")
		partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd.Flags().String("max-results", "", "Use this parameter to control the number of items returned in each request, which can be useful for performance tuning and managing large result sets.")
		partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd.Flags().String("next-token", "", "Use this parameter for pagination when the result set spans multiple pages.")
		partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd.Flags().String("opportunity-identifier", "", "Filters tasks by the identifiers of the opportunities they created or are associated with.")
		partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd.Flags().String("sort", "", "Specifies the sorting criteria for the returned results.")
		partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd.Flags().String("task-identifier", "", "Filters tasks by their unique identifiers.")
		partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd.Flags().String("task-status", "", "Filters the tasks based on their current status.")
		partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd.MarkFlagRequired("catalog")
	})
	partnercentralSellingCmd.AddCommand(partnercentralSelling_listEngagementByAcceptingInvitationTasksCmd)
}
