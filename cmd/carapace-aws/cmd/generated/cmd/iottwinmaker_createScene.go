package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_createSceneCmd = &cobra.Command{
	Use:   "create-scene",
	Short: "Creates a scene.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_createSceneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_createSceneCmd).Standalone()

		iottwinmaker_createSceneCmd.Flags().String("capabilities", "", "A list of capabilities that the scene uses to render itself.")
		iottwinmaker_createSceneCmd.Flags().String("content-location", "", "The relative path that specifies the location of the content definition file.")
		iottwinmaker_createSceneCmd.Flags().String("description", "", "The description for this scene.")
		iottwinmaker_createSceneCmd.Flags().String("scene-id", "", "The ID of the scene.")
		iottwinmaker_createSceneCmd.Flags().String("scene-metadata", "", "The request metadata.")
		iottwinmaker_createSceneCmd.Flags().String("tags", "", "Metadata that you can use to manage the scene.")
		iottwinmaker_createSceneCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the scene.")
		iottwinmaker_createSceneCmd.MarkFlagRequired("content-location")
		iottwinmaker_createSceneCmd.MarkFlagRequired("scene-id")
		iottwinmaker_createSceneCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_createSceneCmd)
}
