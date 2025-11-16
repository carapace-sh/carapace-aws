package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_updateApprovalTeamCmd = &cobra.Command{
	Use:   "update-approval-team",
	Short: "Updates an approval team.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_updateApprovalTeamCmd).Standalone()

	mpa_updateApprovalTeamCmd.Flags().String("approval-strategy", "", "An `ApprovalStrategy` object.")
	mpa_updateApprovalTeamCmd.Flags().String("approvers", "", "An array of `ApprovalTeamRequestApprover` objects.")
	mpa_updateApprovalTeamCmd.Flags().String("arn", "", "Amazon Resource Name (ARN) for the team.")
	mpa_updateApprovalTeamCmd.Flags().String("description", "", "Description for the team.")
	mpa_updateApprovalTeamCmd.MarkFlagRequired("arn")
	mpaCmd.AddCommand(mpa_updateApprovalTeamCmd)
}
