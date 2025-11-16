package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_createCustomActionTypeCmd = &cobra.Command{
	Use:   "create-custom-action-type",
	Short: "Creates a new custom action that can be used in all pipelines associated with the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_createCustomActionTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_createCustomActionTypeCmd).Standalone()

		codepipeline_createCustomActionTypeCmd.Flags().String("category", "", "The category of the custom action, such as a build action or a test action.")
		codepipeline_createCustomActionTypeCmd.Flags().String("configuration-properties", "", "The configuration properties for the custom action.")
		codepipeline_createCustomActionTypeCmd.Flags().String("input-artifact-details", "", "The details of the input artifact for the action, such as its commit ID.")
		codepipeline_createCustomActionTypeCmd.Flags().String("output-artifact-details", "", "The details of the output artifact of the action, such as its commit ID.")
		codepipeline_createCustomActionTypeCmd.Flags().String("provider", "", "The provider of the service used in the custom action, such as CodeDeploy.")
		codepipeline_createCustomActionTypeCmd.Flags().String("settings", "", "URLs that provide users information about this custom action.")
		codepipeline_createCustomActionTypeCmd.Flags().String("tags", "", "The tags for the custom action.")
		codepipeline_createCustomActionTypeCmd.Flags().String("version", "", "The version identifier of the custom action.")
		codepipeline_createCustomActionTypeCmd.MarkFlagRequired("category")
		codepipeline_createCustomActionTypeCmd.MarkFlagRequired("input-artifact-details")
		codepipeline_createCustomActionTypeCmd.MarkFlagRequired("output-artifact-details")
		codepipeline_createCustomActionTypeCmd.MarkFlagRequired("provider")
		codepipeline_createCustomActionTypeCmd.MarkFlagRequired("version")
	})
	codepipelineCmd.AddCommand(codepipeline_createCustomActionTypeCmd)
}
