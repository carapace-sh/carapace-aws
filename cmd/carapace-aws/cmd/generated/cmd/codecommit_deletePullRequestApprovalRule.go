package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_deletePullRequestApprovalRuleCmd = &cobra.Command{
	Use:   "delete-pull-request-approval-rule",
	Short: "Deletes an approval rule from a specified pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_deletePullRequestApprovalRuleCmd).Standalone()

	codecommit_deletePullRequestApprovalRuleCmd.Flags().String("approval-rule-name", "", "The name of the approval rule you want to delete.")
	codecommit_deletePullRequestApprovalRuleCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request that contains the approval rule you want to delete.")
	codecommit_deletePullRequestApprovalRuleCmd.MarkFlagRequired("approval-rule-name")
	codecommit_deletePullRequestApprovalRuleCmd.MarkFlagRequired("pull-request-id")
	codecommitCmd.AddCommand(codecommit_deletePullRequestApprovalRuleCmd)
}
