package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_listApprovalRuleTemplatesCmd = &cobra.Command{
	Use:   "list-approval-rule-templates",
	Short: "Lists all approval rule templates in the specified Amazon Web Services Region in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_listApprovalRuleTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecommit_listApprovalRuleTemplatesCmd).Standalone()

		codecommit_listApprovalRuleTemplatesCmd.Flags().String("max-results", "", "A non-zero, non-negative integer used to limit the number of returned results.")
		codecommit_listApprovalRuleTemplatesCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
	})
	codecommitCmd.AddCommand(codecommit_listApprovalRuleTemplatesCmd)
}
