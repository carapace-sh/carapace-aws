package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updateApprovalRuleTemplateContentCmd = &cobra.Command{
	Use:   "update-approval-rule-template-content",
	Short: "Updates the content of an approval rule template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updateApprovalRuleTemplateContentCmd).Standalone()

	codecommit_updateApprovalRuleTemplateContentCmd.Flags().String("approval-rule-template-name", "", "The name of the approval rule template where you want to update the content of the rule.")
	codecommit_updateApprovalRuleTemplateContentCmd.Flags().String("existing-rule-content-sha256", "", "The SHA-256 hash signature for the content of the approval rule.")
	codecommit_updateApprovalRuleTemplateContentCmd.Flags().String("new-rule-content", "", "The content that replaces the existing content of the rule.")
	codecommit_updateApprovalRuleTemplateContentCmd.MarkFlagRequired("approval-rule-template-name")
	codecommit_updateApprovalRuleTemplateContentCmd.MarkFlagRequired("new-rule-content")
	codecommitCmd.AddCommand(codecommit_updateApprovalRuleTemplateContentCmd)
}
