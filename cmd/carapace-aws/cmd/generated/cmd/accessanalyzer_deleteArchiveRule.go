package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_deleteArchiveRuleCmd = &cobra.Command{
	Use:   "delete-archive-rule",
	Short: "Deletes the specified archive rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_deleteArchiveRuleCmd).Standalone()

	accessanalyzer_deleteArchiveRuleCmd.Flags().String("analyzer-name", "", "The name of the analyzer that associated with the archive rule to delete.")
	accessanalyzer_deleteArchiveRuleCmd.Flags().String("client-token", "", "A client token.")
	accessanalyzer_deleteArchiveRuleCmd.Flags().String("rule-name", "", "The name of the rule to delete.")
	accessanalyzer_deleteArchiveRuleCmd.MarkFlagRequired("analyzer-name")
	accessanalyzer_deleteArchiveRuleCmd.MarkFlagRequired("rule-name")
	accessanalyzerCmd.AddCommand(accessanalyzer_deleteArchiveRuleCmd)
}
