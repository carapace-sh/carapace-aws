package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_getApprovalRuleTemplateCmd = &cobra.Command{
	Use:   "get-approval-rule-template",
	Short: "Returns information about a specified approval rule template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_getApprovalRuleTemplateCmd).Standalone()

	codecommit_getApprovalRuleTemplateCmd.Flags().String("approval-rule-template-name", "", "The name of the approval rule template for which you want to get information.")
	codecommit_getApprovalRuleTemplateCmd.MarkFlagRequired("approval-rule-template-name")
	codecommitCmd.AddCommand(codecommit_getApprovalRuleTemplateCmd)
}
