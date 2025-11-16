package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_listArchiveRulesCmd = &cobra.Command{
	Use:   "list-archive-rules",
	Short: "Retrieves a list of archive rules created for the specified analyzer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_listArchiveRulesCmd).Standalone()

	accessanalyzer_listArchiveRulesCmd.Flags().String("analyzer-name", "", "The name of the analyzer to retrieve rules from.")
	accessanalyzer_listArchiveRulesCmd.Flags().String("max-results", "", "The maximum number of results to return in the request.")
	accessanalyzer_listArchiveRulesCmd.Flags().String("next-token", "", "A token used for pagination of results returned.")
	accessanalyzer_listArchiveRulesCmd.MarkFlagRequired("analyzer-name")
	accessanalyzerCmd.AddCommand(accessanalyzer_listArchiveRulesCmd)
}
