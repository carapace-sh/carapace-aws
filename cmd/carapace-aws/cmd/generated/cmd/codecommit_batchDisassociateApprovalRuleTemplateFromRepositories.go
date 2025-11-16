package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_batchDisassociateApprovalRuleTemplateFromRepositoriesCmd = &cobra.Command{
	Use:   "batch-disassociate-approval-rule-template-from-repositories",
	Short: "Removes the association between an approval rule template and one or more specified repositories.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_batchDisassociateApprovalRuleTemplateFromRepositoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_batchDisassociateApprovalRuleTemplateFromRepositoriesCmd).Standalone()

		codecommit_batchDisassociateApprovalRuleTemplateFromRepositoriesCmd.Flags().String("approval-rule-template-name", "", "The name of the template that you want to disassociate from one or more repositories.")
		codecommit_batchDisassociateApprovalRuleTemplateFromRepositoriesCmd.Flags().String("repository-names", "", "The repository names that you want to disassociate from the approval rule template.")
		codecommit_batchDisassociateApprovalRuleTemplateFromRepositoriesCmd.MarkFlagRequired("approval-rule-template-name")
		codecommit_batchDisassociateApprovalRuleTemplateFromRepositoriesCmd.MarkFlagRequired("repository-names")
	})
	codecommitCmd.AddCommand(codecommit_batchDisassociateApprovalRuleTemplateFromRepositoriesCmd)
}
