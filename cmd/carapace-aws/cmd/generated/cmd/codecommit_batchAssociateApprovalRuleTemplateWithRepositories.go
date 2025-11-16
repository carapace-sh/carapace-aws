package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_batchAssociateApprovalRuleTemplateWithRepositoriesCmd = &cobra.Command{
	Use:   "batch-associate-approval-rule-template-with-repositories",
	Short: "Creates an association between an approval rule template and one or more specified repositories.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_batchAssociateApprovalRuleTemplateWithRepositoriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_batchAssociateApprovalRuleTemplateWithRepositoriesCmd).Standalone()

		codecommit_batchAssociateApprovalRuleTemplateWithRepositoriesCmd.Flags().String("approval-rule-template-name", "", "The name of the template you want to associate with one or more repositories.")
		codecommit_batchAssociateApprovalRuleTemplateWithRepositoriesCmd.Flags().String("repository-names", "", "The names of the repositories you want to associate with the template.")
		codecommit_batchAssociateApprovalRuleTemplateWithRepositoriesCmd.MarkFlagRequired("approval-rule-template-name")
		codecommit_batchAssociateApprovalRuleTemplateWithRepositoriesCmd.MarkFlagRequired("repository-names")
	})
	codecommitCmd.AddCommand(codecommit_batchAssociateApprovalRuleTemplateWithRepositoriesCmd)
}
