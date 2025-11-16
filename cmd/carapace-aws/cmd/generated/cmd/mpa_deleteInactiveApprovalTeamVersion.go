package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_deleteInactiveApprovalTeamVersionCmd = &cobra.Command{
	Use:   "delete-inactive-approval-team-version",
	Short: "Deletes an inactive approval team.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_deleteInactiveApprovalTeamVersionCmd).Standalone()

	mpa_deleteInactiveApprovalTeamVersionCmd.Flags().String("arn", "", "Amaazon Resource Name (ARN) for the team.")
	mpa_deleteInactiveApprovalTeamVersionCmd.Flags().String("version-id", "", "Version ID for the team.")
	mpa_deleteInactiveApprovalTeamVersionCmd.MarkFlagRequired("arn")
	mpa_deleteInactiveApprovalTeamVersionCmd.MarkFlagRequired("version-id")
	mpaCmd.AddCommand(mpa_deleteInactiveApprovalTeamVersionCmd)
}
