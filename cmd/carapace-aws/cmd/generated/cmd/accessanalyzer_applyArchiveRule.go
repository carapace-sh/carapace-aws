package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_applyArchiveRuleCmd = &cobra.Command{
	Use:   "apply-archive-rule",
	Short: "Retroactively applies the archive rule to existing findings that meet the archive rule criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_applyArchiveRuleCmd).Standalone()

	accessanalyzer_applyArchiveRuleCmd.Flags().String("analyzer-arn", "", "The Amazon resource name (ARN) of the analyzer.")
	accessanalyzer_applyArchiveRuleCmd.Flags().String("client-token", "", "A client token.")
	accessanalyzer_applyArchiveRuleCmd.Flags().String("rule-name", "", "The name of the rule to apply.")
	accessanalyzer_applyArchiveRuleCmd.MarkFlagRequired("analyzer-arn")
	accessanalyzer_applyArchiveRuleCmd.MarkFlagRequired("rule-name")
	accessanalyzerCmd.AddCommand(accessanalyzer_applyArchiveRuleCmd)
}
