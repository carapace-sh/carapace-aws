package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_deleteAnalysisTemplateCmd = &cobra.Command{
	Use:   "delete-analysis-template",
	Short: "Deletes an analysis template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_deleteAnalysisTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_deleteAnalysisTemplateCmd).Standalone()

		cleanrooms_deleteAnalysisTemplateCmd.Flags().String("analysis-template-identifier", "", "The identifier for the analysis template resource.")
		cleanrooms_deleteAnalysisTemplateCmd.Flags().String("membership-identifier", "", "The identifier for a membership resource.")
		cleanrooms_deleteAnalysisTemplateCmd.MarkFlagRequired("analysis-template-identifier")
		cleanrooms_deleteAnalysisTemplateCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_deleteAnalysisTemplateCmd)
}
