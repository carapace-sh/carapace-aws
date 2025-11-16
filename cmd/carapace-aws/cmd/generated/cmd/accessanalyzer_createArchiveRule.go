package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_createArchiveRuleCmd = &cobra.Command{
	Use:   "create-archive-rule",
	Short: "Creates an archive rule for the specified analyzer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_createArchiveRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_createArchiveRuleCmd).Standalone()

		accessanalyzer_createArchiveRuleCmd.Flags().String("analyzer-name", "", "The name of the created analyzer.")
		accessanalyzer_createArchiveRuleCmd.Flags().String("client-token", "", "A client token.")
		accessanalyzer_createArchiveRuleCmd.Flags().String("filter", "", "The criteria for the rule.")
		accessanalyzer_createArchiveRuleCmd.Flags().String("rule-name", "", "The name of the rule to create.")
		accessanalyzer_createArchiveRuleCmd.MarkFlagRequired("analyzer-name")
		accessanalyzer_createArchiveRuleCmd.MarkFlagRequired("filter")
		accessanalyzer_createArchiveRuleCmd.MarkFlagRequired("rule-name")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_createArchiveRuleCmd)
}
