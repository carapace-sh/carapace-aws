package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_getArchiveRuleCmd = &cobra.Command{
	Use:   "get-archive-rule",
	Short: "Retrieves information about an archive rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_getArchiveRuleCmd).Standalone()

	accessanalyzer_getArchiveRuleCmd.Flags().String("analyzer-name", "", "The name of the analyzer to retrieve rules from.")
	accessanalyzer_getArchiveRuleCmd.Flags().String("rule-name", "", "The name of the rule to retrieve.")
	accessanalyzer_getArchiveRuleCmd.MarkFlagRequired("analyzer-name")
	accessanalyzer_getArchiveRuleCmd.MarkFlagRequired("rule-name")
	accessanalyzerCmd.AddCommand(accessanalyzer_getArchiveRuleCmd)
}
