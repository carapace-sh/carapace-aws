package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var detective_rejectInvitationCmd = &cobra.Command{
	Use:   "reject-invitation",
	Short: "Rejects an invitation to contribute the account data to a behavior graph.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(detective_rejectInvitationCmd).Standalone()

	detective_rejectInvitationCmd.Flags().String("graph-arn", "", "The ARN of the behavior graph to reject the invitation to.")
	detective_rejectInvitationCmd.MarkFlagRequired("graph-arn")
	detectiveCmd.AddCommand(detective_rejectInvitationCmd)
}
