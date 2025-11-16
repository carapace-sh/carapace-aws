package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_createAnalysisTemplateCmd = &cobra.Command{
	Use:   "create-analysis-template",
	Short: "Creates a new analysis template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_createAnalysisTemplateCmd).Standalone()

	cleanrooms_createAnalysisTemplateCmd.Flags().String("analysis-parameters", "", "The parameters of the analysis template.")
	cleanrooms_createAnalysisTemplateCmd.Flags().String("description", "", "The description of the analysis template.")
	cleanrooms_createAnalysisTemplateCmd.Flags().String("error-message-configuration", "", "The configuration that specifies the level of detail in error messages returned by analyses using this template.")
	cleanrooms_createAnalysisTemplateCmd.Flags().String("format", "", "The format of the analysis template.")
	cleanrooms_createAnalysisTemplateCmd.Flags().String("membership-identifier", "", "The identifier for a membership resource.")
	cleanrooms_createAnalysisTemplateCmd.Flags().String("name", "", "The name of the analysis template.")
	cleanrooms_createAnalysisTemplateCmd.Flags().String("schema", "", "")
	cleanrooms_createAnalysisTemplateCmd.Flags().String("source", "", "The information in the analysis template.")
	cleanrooms_createAnalysisTemplateCmd.Flags().String("tags", "", "An optional label that you can assign to a resource when you create it.")
	cleanrooms_createAnalysisTemplateCmd.MarkFlagRequired("format")
	cleanrooms_createAnalysisTemplateCmd.MarkFlagRequired("membership-identifier")
	cleanrooms_createAnalysisTemplateCmd.MarkFlagRequired("name")
	cleanrooms_createAnalysisTemplateCmd.MarkFlagRequired("source")
	cleanroomsCmd.AddCommand(cleanrooms_createAnalysisTemplateCmd)
}
