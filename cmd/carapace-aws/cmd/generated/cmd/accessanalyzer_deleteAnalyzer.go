package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_deleteAnalyzerCmd = &cobra.Command{
	Use:   "delete-analyzer",
	Short: "Deletes the specified analyzer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_deleteAnalyzerCmd).Standalone()

	accessanalyzer_deleteAnalyzerCmd.Flags().String("analyzer-name", "", "The name of the analyzer to delete.")
	accessanalyzer_deleteAnalyzerCmd.Flags().String("client-token", "", "A client token.")
	accessanalyzer_deleteAnalyzerCmd.MarkFlagRequired("analyzer-name")
	accessanalyzerCmd.AddCommand(accessanalyzer_deleteAnalyzerCmd)
}
