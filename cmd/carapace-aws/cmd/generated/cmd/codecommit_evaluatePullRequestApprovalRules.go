package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_evaluatePullRequestApprovalRulesCmd = &cobra.Command{
	Use:   "evaluate-pull-request-approval-rules",
	Short: "Evaluates whether a pull request has met all the conditions specified in its associated approval rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_evaluatePullRequestApprovalRulesCmd).Standalone()

	codecommit_evaluatePullRequestApprovalRulesCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request you want to evaluate.")
	codecommit_evaluatePullRequestApprovalRulesCmd.Flags().String("revision-id", "", "The system-generated ID for the pull request revision.")
	codecommit_evaluatePullRequestApprovalRulesCmd.MarkFlagRequired("pull-request-id")
	codecommit_evaluatePullRequestApprovalRulesCmd.MarkFlagRequired("revision-id")
	codecommitCmd.AddCommand(codecommit_evaluatePullRequestApprovalRulesCmd)
}
