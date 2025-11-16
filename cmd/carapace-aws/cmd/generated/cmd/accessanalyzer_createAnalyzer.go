package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_createAnalyzerCmd = &cobra.Command{
	Use:   "create-analyzer",
	Short: "Creates an analyzer for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_createAnalyzerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_createAnalyzerCmd).Standalone()

		accessanalyzer_createAnalyzerCmd.Flags().String("analyzer-name", "", "The name of the analyzer to create.")
		accessanalyzer_createAnalyzerCmd.Flags().String("archive-rules", "", "Specifies the archive rules to add for the analyzer.")
		accessanalyzer_createAnalyzerCmd.Flags().String("client-token", "", "A client token.")
		accessanalyzer_createAnalyzerCmd.Flags().String("configuration", "", "Specifies the configuration of the analyzer.")
		accessanalyzer_createAnalyzerCmd.Flags().String("tags", "", "An array of key-value pairs to apply to the analyzer.")
		accessanalyzer_createAnalyzerCmd.Flags().String("type", "", "The type of analyzer to create.")
		accessanalyzer_createAnalyzerCmd.MarkFlagRequired("analyzer-name")
		accessanalyzer_createAnalyzerCmd.MarkFlagRequired("type")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_createAnalyzerCmd)
}
