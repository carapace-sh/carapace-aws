package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listCollaborationChangeRequestsCmd = &cobra.Command{
	Use:   "list-collaboration-change-requests",
	Short: "Lists all change requests for a collaboration with pagination support.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listCollaborationChangeRequestsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_listCollaborationChangeRequestsCmd).Standalone()

		cleanrooms_listCollaborationChangeRequestsCmd.Flags().String("collaboration-identifier", "", "The identifier of the collaboration that the change request is made against.")
		cleanrooms_listCollaborationChangeRequestsCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
		cleanrooms_listCollaborationChangeRequestsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		cleanrooms_listCollaborationChangeRequestsCmd.Flags().String("status", "", "A filter to only return change requests with the specified status.")
		cleanrooms_listCollaborationChangeRequestsCmd.MarkFlagRequired("collaboration-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_listCollaborationChangeRequestsCmd)
}
