package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_associateApprovalRuleTemplateWithRepositoryCmd = &cobra.Command{
	Use:   "associate-approval-rule-template-with-repository",
	Short: "Creates an association between an approval rule template and a specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_associateApprovalRuleTemplateWithRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_associateApprovalRuleTemplateWithRepositoryCmd).Standalone()

		codecommit_associateApprovalRuleTemplateWithRepositoryCmd.Flags().String("approval-rule-template-name", "", "The name for the approval rule template.")
		codecommit_associateApprovalRuleTemplateWithRepositoryCmd.Flags().String("repository-name", "", "The name of the repository that you want to associate with the template.")
		codecommit_associateApprovalRuleTemplateWithRepositoryCmd.MarkFlagRequired("approval-rule-template-name")
		codecommit_associateApprovalRuleTemplateWithRepositoryCmd.MarkFlagRequired("repository-name")
	})
	codecommitCmd.AddCommand(codecommit_associateApprovalRuleTemplateWithRepositoryCmd)
}
