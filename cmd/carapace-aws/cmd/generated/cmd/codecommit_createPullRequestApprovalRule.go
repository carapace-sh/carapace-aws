package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_createPullRequestApprovalRuleCmd = &cobra.Command{
	Use:   "create-pull-request-approval-rule",
	Short: "Creates an approval rule for a pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_createPullRequestApprovalRuleCmd).Standalone()

	codecommit_createPullRequestApprovalRuleCmd.Flags().String("approval-rule-content", "", "The content of the approval rule, including the number of approvals needed and the structure of an approval pool defined for approvals, if any.")
	codecommit_createPullRequestApprovalRuleCmd.Flags().String("approval-rule-name", "", "The name for the approval rule.")
	codecommit_createPullRequestApprovalRuleCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request for which you want to create the approval rule.")
	codecommit_createPullRequestApprovalRuleCmd.MarkFlagRequired("approval-rule-content")
	codecommit_createPullRequestApprovalRuleCmd.MarkFlagRequired("approval-rule-name")
	codecommit_createPullRequestApprovalRuleCmd.MarkFlagRequired("pull-request-id")
	codecommitCmd.AddCommand(codecommit_createPullRequestApprovalRuleCmd)
}
