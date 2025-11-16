package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updatePullRequestApprovalRuleContentCmd = &cobra.Command{
	Use:   "update-pull-request-approval-rule-content",
	Short: "Updates the structure of an approval rule created specifically for a pull request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updatePullRequestApprovalRuleContentCmd).Standalone()

	codecommit_updatePullRequestApprovalRuleContentCmd.Flags().String("approval-rule-name", "", "The name of the approval rule you want to update.")
	codecommit_updatePullRequestApprovalRuleContentCmd.Flags().String("existing-rule-content-sha256", "", "The SHA-256 hash signature for the content of the approval rule.")
	codecommit_updatePullRequestApprovalRuleContentCmd.Flags().String("new-rule-content", "", "The updated content for the approval rule.")
	codecommit_updatePullRequestApprovalRuleContentCmd.Flags().String("pull-request-id", "", "The system-generated ID of the pull request.")
	codecommit_updatePullRequestApprovalRuleContentCmd.MarkFlagRequired("approval-rule-name")
	codecommit_updatePullRequestApprovalRuleContentCmd.MarkFlagRequired("new-rule-content")
	codecommit_updatePullRequestApprovalRuleContentCmd.MarkFlagRequired("pull-request-id")
	codecommitCmd.AddCommand(codecommit_updatePullRequestApprovalRuleContentCmd)
}
