package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_deleteSceneCmd = &cobra.Command{
	Use:   "delete-scene",
	Short: "Deletes a scene.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_deleteSceneCmd).Standalone()

	iottwinmaker_deleteSceneCmd.Flags().String("scene-id", "", "The ID of the scene to delete.")
	iottwinmaker_deleteSceneCmd.Flags().String("workspace-id", "", "The ID of the workspace.")
	iottwinmaker_deleteSceneCmd.MarkFlagRequired("scene-id")
	iottwinmaker_deleteSceneCmd.MarkFlagRequired("workspace-id")
	iottwinmakerCmd.AddCommand(iottwinmaker_deleteSceneCmd)
}
