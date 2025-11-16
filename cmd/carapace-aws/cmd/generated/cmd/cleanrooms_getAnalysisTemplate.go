package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getAnalysisTemplateCmd = &cobra.Command{
	Use:   "get-analysis-template",
	Short: "Retrieves an analysis template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getAnalysisTemplateCmd).Standalone()

	cleanrooms_getAnalysisTemplateCmd.Flags().String("analysis-template-identifier", "", "The identifier for the analysis template resource.")
	cleanrooms_getAnalysisTemplateCmd.Flags().String("membership-identifier", "", "The identifier for a membership resource.")
	cleanrooms_getAnalysisTemplateCmd.MarkFlagRequired("analysis-template-identifier")
	cleanrooms_getAnalysisTemplateCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_getAnalysisTemplateCmd)
}
