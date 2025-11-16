package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_batchGetCaseRuleCmd = &cobra.Command{
	Use:   "batch-get-case-rule",
	Short: "Gets a batch of case rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_batchGetCaseRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_batchGetCaseRuleCmd).Standalone()

		connectcases_batchGetCaseRuleCmd.Flags().String("case-rules", "", "A list of case rule identifiers.")
		connectcases_batchGetCaseRuleCmd.Flags().String("domain-id", "", "Unique identifier of a Cases domain.")
		connectcases_batchGetCaseRuleCmd.MarkFlagRequired("case-rules")
		connectcases_batchGetCaseRuleCmd.MarkFlagRequired("domain-id")
	})
	connectcasesCmd.AddCommand(connectcases_batchGetCaseRuleCmd)
}
