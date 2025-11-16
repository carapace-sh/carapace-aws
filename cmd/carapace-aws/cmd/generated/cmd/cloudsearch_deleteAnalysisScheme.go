package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_deleteAnalysisSchemeCmd = &cobra.Command{
	Use:   "delete-analysis-scheme",
	Short: "Deletes an analysis scheme.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_deleteAnalysisSchemeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudsearch_deleteAnalysisSchemeCmd).Standalone()

		cloudsearch_deleteAnalysisSchemeCmd.Flags().String("analysis-scheme-name", "", "The name of the analysis scheme you want to delete.")
		cloudsearch_deleteAnalysisSchemeCmd.Flags().String("domain-name", "", "")
		cloudsearch_deleteAnalysisSchemeCmd.MarkFlagRequired("analysis-scheme-name")
		cloudsearch_deleteAnalysisSchemeCmd.MarkFlagRequired("domain-name")
	})
	cloudsearchCmd.AddCommand(cloudsearch_deleteAnalysisSchemeCmd)
}
