package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updatePullRequestApprovalStateCmd = &cobra.Command{
	Use:   "update-pull-request-approval-state",
	Short: "Updates the state of a user's approval on a pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updatePullRequestApprovalStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_updatePullRequestApprovalStateCmd).Standalone()

		codecommit_updatePullRequestApprovalStateCmd.Flags().String("approval-state", "", "The approval state to associate with the user on the pull request.")
		codecommit_updatePullRequestApprovalStateCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request.")
		codecommit_updatePullRequestApprovalStateCmd.Flags().String("revision-id", "", "The system-generated ID of the revision.")
		codecommit_updatePullRequestApprovalStateCmd.MarkFlagRequired("approval-state")
		codecommit_updatePullRequestApprovalStateCmd.MarkFlagRequired("pull-request-id")
		codecommit_updatePullRequestApprovalStateCmd.MarkFlagRequired("revision-id")
	})
	codecommitCmd.AddCommand(codecommit_updatePullRequestApprovalStateCmd)
}
