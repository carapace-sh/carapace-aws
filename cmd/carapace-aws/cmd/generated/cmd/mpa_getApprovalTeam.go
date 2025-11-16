package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_getApprovalTeamCmd = &cobra.Command{
	Use:   "get-approval-team",
	Short: "Returns details for an approval team.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_getApprovalTeamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_getApprovalTeamCmd).Standalone()

		mpa_getApprovalTeamCmd.Flags().String("arn", "", "Amazon Resource Name (ARN) for the team.")
		mpa_getApprovalTeamCmd.MarkFlagRequired("arn")
	})
	mpaCmd.AddCommand(mpa_getApprovalTeamCmd)
}
