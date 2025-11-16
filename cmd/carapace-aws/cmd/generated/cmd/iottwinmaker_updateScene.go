package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_updateSceneCmd = &cobra.Command{
	Use:   "update-scene",
	Short: "Updates a scene.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_updateSceneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_updateSceneCmd).Standalone()

		iottwinmaker_updateSceneCmd.Flags().String("capabilities", "", "A list of capabilities that the scene uses to render.")
		iottwinmaker_updateSceneCmd.Flags().String("content-location", "", "The relative path that specifies the location of the content definition file.")
		iottwinmaker_updateSceneCmd.Flags().String("description", "", "The description of this scene.")
		iottwinmaker_updateSceneCmd.Flags().String("scene-id", "", "The ID of the scene.")
		iottwinmaker_updateSceneCmd.Flags().String("scene-metadata", "", "The scene metadata.")
		iottwinmaker_updateSceneCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the scene.")
		iottwinmaker_updateSceneCmd.MarkFlagRequired("scene-id")
		iottwinmaker_updateSceneCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_updateSceneCmd)
}
