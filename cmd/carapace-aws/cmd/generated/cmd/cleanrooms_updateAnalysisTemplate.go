package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updateAnalysisTemplateCmd = &cobra.Command{
	Use:   "update-analysis-template",
	Short: "Updates the analysis template metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updateAnalysisTemplateCmd).Standalone()

	cleanrooms_updateAnalysisTemplateCmd.Flags().String("analysis-template-identifier", "", "The identifier for the analysis template resource.")
	cleanrooms_updateAnalysisTemplateCmd.Flags().String("description", "", "A new description for the analysis template.")
	cleanrooms_updateAnalysisTemplateCmd.Flags().String("membership-identifier", "", "The identifier for a membership resource.")
	cleanrooms_updateAnalysisTemplateCmd.MarkFlagRequired("analysis-template-identifier")
	cleanrooms_updateAnalysisTemplateCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_updateAnalysisTemplateCmd)
}
