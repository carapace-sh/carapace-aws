package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_listInvitationsCmd = &cobra.Command{
	Use:   "list-invitations",
	Short: "Retrieves the list of open and accepted behavior graph invitations for the member account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_listInvitationsCmd).Standalone()

	detective_listInvitationsCmd.Flags().String("max-results", "", "The maximum number of behavior graph invitations to return in the response.")
	detective_listInvitationsCmd.Flags().String("next-token", "", "For requests to retrieve the next page of results, the pagination token that was returned with the previous page of results.")
	detectiveCmd.AddCommand(detective_listInvitationsCmd)
}
