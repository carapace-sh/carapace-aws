package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_listAssociatedApprovalRuleTemplatesForRepositoryCmd = &cobra.Command{
	Use:   "list-associated-approval-rule-templates-for-repository",
	Short: "Lists all approval rule templates that are associated with a specified repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_listAssociatedApprovalRuleTemplatesForRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_listAssociatedApprovalRuleTemplatesForRepositoryCmd).Standalone()

		codecommit_listAssociatedApprovalRuleTemplatesForRepositoryCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
		codecommit_listAssociatedApprovalRuleTemplatesForRepositoryCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
		codecommit_listAssociatedApprovalRuleTemplatesForRepositoryCmd.Flags().String("repository-name", "", "The name of the repository for which you want to list all associated approval rule templates.")
		codecommit_listAssociatedApprovalRuleTemplatesForRepositoryCmd.MarkFlagRequired("repository-name")
	})
	codecommitCmd.AddCommand(codecommit_listAssociatedApprovalRuleTemplatesForRepositoryCmd)
}
