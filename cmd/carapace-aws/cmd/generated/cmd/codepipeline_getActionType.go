package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_getActionTypeCmd = &cobra.Command{
	Use:   "get-action-type",
	Short: "Returns information about an action type created for an external provider, where the action is to be used by customers of the external provider.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_getActionTypeCmd).Standalone()

	codepipeline_getActionTypeCmd.Flags().String("category", "", "Defines what kind of action can be taken in the stage.")
	codepipeline_getActionTypeCmd.Flags().String("owner", "", "The creator of an action type that was created with any supported integration model.")
	codepipeline_getActionTypeCmd.Flags().String("provider", "", "The provider of the action type being called.")
	codepipeline_getActionTypeCmd.Flags().String("version", "", "A string that describes the action type version.")
	codepipeline_getActionTypeCmd.MarkFlagRequired("category")
	codepipeline_getActionTypeCmd.MarkFlagRequired("owner")
	codepipeline_getActionTypeCmd.MarkFlagRequired("provider")
	codepipeline_getActionTypeCmd.MarkFlagRequired("version")
	codepipelineCmd.AddCommand(codepipeline_getActionTypeCmd)
}
