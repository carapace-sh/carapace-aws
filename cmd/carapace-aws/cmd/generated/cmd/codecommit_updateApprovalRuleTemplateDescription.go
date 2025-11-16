package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_updateApprovalRuleTemplateDescriptionCmd = &cobra.Command{
	Use:   "update-approval-rule-template-description",
	Short: "Updates the description for a specified approval rule template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_updateApprovalRuleTemplateDescriptionCmd).Standalone()

	codecommit_updateApprovalRuleTemplateDescriptionCmd.Flags().String("approval-rule-template-description", "", "The updated description of the approval rule template.")
	codecommit_updateApprovalRuleTemplateDescriptionCmd.Flags().String("approval-rule-template-name", "", "The name of the template for which you want to update the description.")
	codecommit_updateApprovalRuleTemplateDescriptionCmd.MarkFlagRequired("approval-rule-template-description")
	codecommit_updateApprovalRuleTemplateDescriptionCmd.MarkFlagRequired("approval-rule-template-name")
	codecommitCmd.AddCommand(codecommit_updateApprovalRuleTemplateDescriptionCmd)
}
