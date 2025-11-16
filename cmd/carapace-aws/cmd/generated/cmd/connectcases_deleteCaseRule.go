package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_deleteCaseRuleCmd = &cobra.Command{
	Use:   "delete-case-rule",
	Short: "Deletes a case rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_deleteCaseRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_deleteCaseRuleCmd).Standalone()

		connectcases_deleteCaseRuleCmd.Flags().String("case-rule-id", "", "Unique identifier of a case rule.")
		connectcases_deleteCaseRuleCmd.Flags().String("domain-id", "", "Unique identifier of a Cases domain.")
		connectcases_deleteCaseRuleCmd.MarkFlagRequired("case-rule-id")
		connectcases_deleteCaseRuleCmd.MarkFlagRequired("domain-id")
	})
	connectcasesCmd.AddCommand(connectcases_deleteCaseRuleCmd)
}
