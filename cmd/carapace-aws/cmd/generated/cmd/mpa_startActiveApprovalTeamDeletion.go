package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_startActiveApprovalTeamDeletionCmd = &cobra.Command{
	Use:   "start-active-approval-team-deletion",
	Short: "Starts the deletion process for an active approval team.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_startActiveApprovalTeamDeletionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_startActiveApprovalTeamDeletionCmd).Standalone()

		mpa_startActiveApprovalTeamDeletionCmd.Flags().String("arn", "", "Amazon Resource Name (ARN) for the team.")
		mpa_startActiveApprovalTeamDeletionCmd.Flags().String("pending-window-days", "", "Number of days between when the team approves the delete request and when the team is deleted.")
		mpa_startActiveApprovalTeamDeletionCmd.MarkFlagRequired("arn")
	})
	mpaCmd.AddCommand(mpa_startActiveApprovalTeamDeletionCmd)
}
