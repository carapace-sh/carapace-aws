package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_deleteCustomActionTypeCmd = &cobra.Command{
	Use:   "delete-custom-action-type",
	Short: "Marks a custom action as deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_deleteCustomActionTypeCmd).Standalone()

	codepipeline_deleteCustomActionTypeCmd.Flags().String("category", "", "The category of the custom action that you want to delete, such as source or deploy.")
	codepipeline_deleteCustomActionTypeCmd.Flags().String("provider", "", "The provider of the service used in the custom action, such as CodeDeploy.")
	codepipeline_deleteCustomActionTypeCmd.Flags().String("version", "", "The version of the custom action to delete.")
	codepipeline_deleteCustomActionTypeCmd.MarkFlagRequired("category")
	codepipeline_deleteCustomActionTypeCmd.MarkFlagRequired("provider")
	codepipeline_deleteCustomActionTypeCmd.MarkFlagRequired("version")
	codepipelineCmd.AddCommand(codepipeline_deleteCustomActionTypeCmd)
}
