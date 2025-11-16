package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getCollaborationAnalysisTemplateCmd = &cobra.Command{
	Use:   "get-collaboration-analysis-template",
	Short: "Retrieves an analysis template within a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getCollaborationAnalysisTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_getCollaborationAnalysisTemplateCmd).Standalone()

		cleanrooms_getCollaborationAnalysisTemplateCmd.Flags().String("analysis-template-arn", "", "The Amazon Resource Name (ARN) associated with the analysis template within a collaboration.")
		cleanrooms_getCollaborationAnalysisTemplateCmd.Flags().String("collaboration-identifier", "", "A unique identifier for the collaboration that the analysis templates belong to.")
		cleanrooms_getCollaborationAnalysisTemplateCmd.MarkFlagRequired("analysis-template-arn")
		cleanrooms_getCollaborationAnalysisTemplateCmd.MarkFlagRequired("collaboration-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_getCollaborationAnalysisTemplateCmd)
}
