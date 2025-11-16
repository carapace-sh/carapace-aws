package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_createApprovalTeamCmd = &cobra.Command{
	Use:   "create-approval-team",
	Short: "Creates a new approval team.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_createApprovalTeamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_createApprovalTeamCmd).Standalone()

		mpa_createApprovalTeamCmd.Flags().String("approval-strategy", "", "An `ApprovalStrategy` object.")
		mpa_createApprovalTeamCmd.Flags().String("approvers", "", "An array of `ApprovalTeamRequesterApprovers` objects.")
		mpa_createApprovalTeamCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		mpa_createApprovalTeamCmd.Flags().String("description", "", "Description for the team.")
		mpa_createApprovalTeamCmd.Flags().String("name", "", "Name of the team.")
		mpa_createApprovalTeamCmd.Flags().String("policies", "", "An array of `PolicyReference` objects.")
		mpa_createApprovalTeamCmd.Flags().String("tags", "", "Tags you want to attach to the team.")
		mpa_createApprovalTeamCmd.MarkFlagRequired("approval-strategy")
		mpa_createApprovalTeamCmd.MarkFlagRequired("approvers")
		mpa_createApprovalTeamCmd.MarkFlagRequired("description")
		mpa_createApprovalTeamCmd.MarkFlagRequired("name")
		mpa_createApprovalTeamCmd.MarkFlagRequired("policies")
	})
	mpaCmd.AddCommand(mpa_createApprovalTeamCmd)
}
