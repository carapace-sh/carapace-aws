package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listInvitationsCmd = &cobra.Command{
	Use:   "list-invitations",
	Short: "Retrieves information about Amazon Macie membership invitations that were received by an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listInvitationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_listInvitationsCmd).Standalone()

		macie2_listInvitationsCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of a paginated response.")
		macie2_listInvitationsCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	})
	macie2Cmd.AddCommand(macie2_listInvitationsCmd)
}
