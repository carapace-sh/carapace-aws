package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_disassociateApprovalRuleTemplateFromRepositoryCmd = &cobra.Command{
	Use:   "disassociate-approval-rule-template-from-repository",
	Short: "Removes the association between a template and a repository so that approval rules based on the template are not automatically created when pull requests are created in the specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_disassociateApprovalRuleTemplateFromRepositoryCmd).Standalone()

	codecommit_disassociateApprovalRuleTemplateFromRepositoryCmd.Flags().String("approval-rule-template-name", "", "The name of the approval rule template to disassociate from a specified repository.")
	codecommit_disassociateApprovalRuleTemplateFromRepositoryCmd.Flags().String("repository-name", "", "The name of the repository you want to disassociate from the template.")
	codecommit_disassociateApprovalRuleTemplateFromRepositoryCmd.MarkFlagRequired("approval-rule-template-name")
	codecommit_disassociateApprovalRuleTemplateFromRepositoryCmd.MarkFlagRequired("repository-name")
	codecommitCmd.AddCommand(codecommit_disassociateApprovalRuleTemplateFromRepositoryCmd)
}
