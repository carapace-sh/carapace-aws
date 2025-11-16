package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_updateArchiveRuleCmd = &cobra.Command{
	Use:   "update-archive-rule",
	Short: "Updates the criteria and values for the specified archive rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_updateArchiveRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_updateArchiveRuleCmd).Standalone()

		accessanalyzer_updateArchiveRuleCmd.Flags().String("analyzer-name", "", "The name of the analyzer to update the archive rules for.")
		accessanalyzer_updateArchiveRuleCmd.Flags().String("client-token", "", "A client token.")
		accessanalyzer_updateArchiveRuleCmd.Flags().String("filter", "", "A filter to match for the rules to update.")
		accessanalyzer_updateArchiveRuleCmd.Flags().String("rule-name", "", "The name of the rule to update.")
		accessanalyzer_updateArchiveRuleCmd.MarkFlagRequired("analyzer-name")
		accessanalyzer_updateArchiveRuleCmd.MarkFlagRequired("filter")
		accessanalyzer_updateArchiveRuleCmd.MarkFlagRequired("rule-name")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_updateArchiveRuleCmd)
}
