package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_getSceneCmd = &cobra.Command{
	Use:   "get-scene",
	Short: "Retrieves information about a scene.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_getSceneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iottwinmaker_getSceneCmd).Standalone()

		iottwinmaker_getSceneCmd.Flags().String("scene-id", "", "The ID of the scene.")
		iottwinmaker_getSceneCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the scene.")
		iottwinmaker_getSceneCmd.MarkFlagRequired("scene-id")
		iottwinmaker_getSceneCmd.MarkFlagRequired("workspace-id")
	})
	iottwinmakerCmd.AddCommand(iottwinmaker_getSceneCmd)
}
