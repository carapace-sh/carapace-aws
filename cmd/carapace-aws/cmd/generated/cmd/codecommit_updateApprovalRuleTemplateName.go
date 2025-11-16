package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updateApprovalRuleTemplateNameCmd = &cobra.Command{
	Use:   "update-approval-rule-template-name",
	Short: "Updates the name of a specified approval rule template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updateApprovalRuleTemplateNameCmd).Standalone()

	codecommit_updateApprovalRuleTemplateNameCmd.Flags().String("new-approval-rule-template-name", "", "The new name you want to apply to the approval rule template.")
	codecommit_updateApprovalRuleTemplateNameCmd.Flags().String("old-approval-rule-template-name", "", "The current name of the approval rule template.")
	codecommit_updateApprovalRuleTemplateNameCmd.MarkFlagRequired("new-approval-rule-template-name")
	codecommit_updateApprovalRuleTemplateNameCmd.MarkFlagRequired("old-approval-rule-template-name")
	codecommitCmd.AddCommand(codecommit_updateApprovalRuleTemplateNameCmd)
}
