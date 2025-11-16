package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accessanalyzer_listAnalyzersCmd = &cobra.Command{
	Use:   "list-analyzers",
	Short: "Retrieves a list of analyzers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accessanalyzer_listAnalyzersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(accessanalyzer_listAnalyzersCmd).Standalone()

		accessanalyzer_listAnalyzersCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		accessanalyzer_listAnalyzersCmd.Flags().String("next-token", "", "A token used for pagination of results returned.")
		accessanalyzer_listAnalyzersCmd.Flags().String("type", "", "The type of analyzer.")
	})
	accessanalyzerCmd.AddCommand(accessanalyzer_listAnalyzersCmd)
}
