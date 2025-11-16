package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_createApprovalRuleTemplateCmd = &cobra.Command{
	Use:   "create-approval-rule-template",
	Short: "Creates a template for approval rules that can then be associated with one or more repositories in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_createApprovalRuleTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_createApprovalRuleTemplateCmd).Standalone()

		codecommit_createApprovalRuleTemplateCmd.Flags().String("approval-rule-template-content", "", "The content of the approval rule that is created on pull requests in associated repositories.")
		codecommit_createApprovalRuleTemplateCmd.Flags().String("approval-rule-template-description", "", "The description of the approval rule template.")
		codecommit_createApprovalRuleTemplateCmd.Flags().String("approval-rule-template-name", "", "The name of the approval rule template.")
		codecommit_createApprovalRuleTemplateCmd.MarkFlagRequired("approval-rule-template-content")
		codecommit_createApprovalRuleTemplateCmd.MarkFlagRequired("approval-rule-template-name")
	})
	codecommitCmd.AddCommand(codecommit_createApprovalRuleTemplateCmd)
}
