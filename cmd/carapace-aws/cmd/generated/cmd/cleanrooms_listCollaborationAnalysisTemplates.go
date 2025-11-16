package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listCollaborationAnalysisTemplatesCmd = &cobra.Command{
	Use:   "list-collaboration-analysis-templates",
	Short: "Lists analysis templates within a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listCollaborationAnalysisTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_listCollaborationAnalysisTemplatesCmd).Standalone()

		cleanrooms_listCollaborationAnalysisTemplatesCmd.Flags().String("collaboration-identifier", "", "A unique identifier for the collaboration that the analysis templates belong to.")
		cleanrooms_listCollaborationAnalysisTemplatesCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
		cleanrooms_listCollaborationAnalysisTemplatesCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		cleanrooms_listCollaborationAnalysisTemplatesCmd.MarkFlagRequired("collaboration-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_listCollaborationAnalysisTemplatesCmd)
}
