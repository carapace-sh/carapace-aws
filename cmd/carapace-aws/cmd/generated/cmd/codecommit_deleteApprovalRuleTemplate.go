package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_deleteApprovalRuleTemplateCmd = &cobra.Command{
	Use:   "delete-approval-rule-template",
	Short: "Deletes a specified approval rule template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_deleteApprovalRuleTemplateCmd).Standalone()

	codecommit_deleteApprovalRuleTemplateCmd.Flags().String("approval-rule-template-name", "", "The name of the approval rule template to delete.")
	codecommit_deleteApprovalRuleTemplateCmd.MarkFlagRequired("approval-rule-template-name")
	codecommitCmd.AddCommand(codecommit_deleteApprovalRuleTemplateCmd)
}
