package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_updateAnalyzerCmd = &cobra.Command{
	Use:   "update-analyzer",
	Short: "Modifies the configuration of an existing analyzer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_updateAnalyzerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_updateAnalyzerCmd).Standalone()

		accessanalyzer_updateAnalyzerCmd.Flags().String("analyzer-name", "", "The name of the analyzer to modify.")
		accessanalyzer_updateAnalyzerCmd.Flags().String("configuration", "", "")
		accessanalyzer_updateAnalyzerCmd.MarkFlagRequired("analyzer-name")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_updateAnalyzerCmd)
}
