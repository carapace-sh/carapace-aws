package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearch_defineAnalysisSchemeCmd = &cobra.Command{
	Use:   "define-analysis-scheme",
	Short: "Configures an analysis scheme that can be applied to a `text` or `text-array` field to define language-specific text processing options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearch_defineAnalysisSchemeCmd).Standalone()

	cloudsearch_defineAnalysisSchemeCmd.Flags().String("analysis-scheme", "", "")
	cloudsearch_defineAnalysisSchemeCmd.Flags().String("domain-name", "", "")
	cloudsearch_defineAnalysisSchemeCmd.MarkFlagRequired("analysis-scheme")
	cloudsearch_defineAnalysisSchemeCmd.MarkFlagRequired("domain-name")
	cloudsearchCmd.AddCommand(cloudsearch_defineAnalysisSchemeCmd)
}
