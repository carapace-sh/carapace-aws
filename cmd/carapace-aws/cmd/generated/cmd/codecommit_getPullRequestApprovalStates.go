package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getPullRequestApprovalStatesCmd = &cobra.Command{
	Use:   "get-pull-request-approval-states",
	Short: "Gets information about the approval states for a specified pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getPullRequestApprovalStatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_getPullRequestApprovalStatesCmd).Standalone()

		codecommit_getPullRequestApprovalStatesCmd.Flags().String("pull-request-id", "", "The system-generated ID for the pull request.")
		codecommit_getPullRequestApprovalStatesCmd.Flags().String("revision-id", "", "The system-generated ID for the pull request revision.")
		codecommit_getPullRequestApprovalStatesCmd.MarkFlagRequired("pull-request-id")
		codecommit_getPullRequestApprovalStatesCmd.MarkFlagRequired("revision-id")
	})
	codecommitCmd.AddCommand(codecommit_getPullRequestApprovalStatesCmd)
}
