package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_acceptInvitationCmd = &cobra.Command{
	Use:   "accept-invitation",
	Short: "Accepts an invitation for the member account to contribute data to a behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_acceptInvitationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(detective_acceptInvitationCmd).Standalone()

		detective_acceptInvitationCmd.Flags().String("graph-arn", "", "The ARN of the behavior graph that the member account is accepting the invitation for.")
		detective_acceptInvitationCmd.MarkFlagRequired("graph-arn")
	})
	detectiveCmd.AddCommand(detective_acceptInvitationCmd)
}
