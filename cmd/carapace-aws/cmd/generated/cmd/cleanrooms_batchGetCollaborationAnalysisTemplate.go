package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_batchGetCollaborationAnalysisTemplateCmd = &cobra.Command{
	Use:   "batch-get-collaboration-analysis-template",
	Short: "Retrieves multiple analysis templates within a collaboration by their Amazon Resource Names (ARNs).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_batchGetCollaborationAnalysisTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_batchGetCollaborationAnalysisTemplateCmd).Standalone()

		cleanrooms_batchGetCollaborationAnalysisTemplateCmd.Flags().String("analysis-template-arns", "", "The Amazon Resource Name (ARN) associated with the analysis template within a collaboration.")
		cleanrooms_batchGetCollaborationAnalysisTemplateCmd.Flags().String("collaboration-identifier", "", "A unique identifier for the collaboration that the analysis templates belong to.")
		cleanrooms_batchGetCollaborationAnalysisTemplateCmd.MarkFlagRequired("analysis-template-arns")
		cleanrooms_batchGetCollaborationAnalysisTemplateCmd.MarkFlagRequired("collaboration-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_batchGetCollaborationAnalysisTemplateCmd)
}
