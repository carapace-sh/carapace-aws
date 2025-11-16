package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_getAnalyzerCmd = &cobra.Command{
	Use:   "get-analyzer",
	Short: "Retrieves information about the specified analyzer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_getAnalyzerCmd).Standalone()

	accessanalyzer_getAnalyzerCmd.Flags().String("analyzer-name", "", "The name of the analyzer retrieved.")
	accessanalyzer_getAnalyzerCmd.MarkFlagRequired("analyzer-name")
	accessanalyzerCmd.AddCommand(accessanalyzer_getAnalyzerCmd)
}
