package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_listRepositoriesForApprovalRuleTemplateCmd = &cobra.Command{
	Use:   "list-repositories-for-approval-rule-template",
	Short: "Lists all repositories associated with the specified approval rule template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_listRepositoriesForApprovalRuleTemplateCmd).Standalone()

	codecommit_listRepositoriesForApprovalRuleTemplateCmd.Flags().String("approval-rule-template-name", "", "The name of the approval rule template for which you want to list repositories that are associated with that template.")
	codecommit_listRepositoriesForApprovalRuleTemplateCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
	codecommit_listRepositoriesForApprovalRuleTemplateCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
	codecommit_listRepositoriesForApprovalRuleTemplateCmd.MarkFlagRequired("approval-rule-template-name")
	codecommitCmd.AddCommand(codecommit_listRepositoriesForApprovalRuleTemplateCmd)
}
