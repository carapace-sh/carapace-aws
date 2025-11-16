package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_listInvitationsCmd = &cobra.Command{
	Use:   "list-invitations",
	Short: "We recommend using Organizations instead of Security Hub invitations to manage your member accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_listInvitationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_listInvitationsCmd).Standalone()

		securityhub_listInvitationsCmd.Flags().String("max-results", "", "The maximum number of items to return in the response.")
		securityhub_listInvitationsCmd.Flags().String("next-token", "", "The token that is required for pagination.")
	})
	securityhubCmd.AddCommand(securityhub_listInvitationsCmd)
}
