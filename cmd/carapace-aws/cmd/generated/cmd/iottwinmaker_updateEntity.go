package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_updateEntityCmd = &cobra.Command{
	Use:   "update-entity",
	Short: "Updates an entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_updateEntityCmd).Standalone()

	iottwinmaker_updateEntityCmd.Flags().String("component-updates", "", "An object that maps strings to the component updates in the request.")
	iottwinmaker_updateEntityCmd.Flags().String("composite-component-updates", "", "This is an object that maps strings to `compositeComponent` updates in the request.")
	iottwinmaker_updateEntityCmd.Flags().String("description", "", "The description of the entity.")
	iottwinmaker_updateEntityCmd.Flags().String("entity-id", "", "The ID of the entity.")
	iottwinmaker_updateEntityCmd.Flags().String("entity-name", "", "The name of the entity.")
	iottwinmaker_updateEntityCmd.Flags().String("parent-entity-update", "", "An object that describes the update request for a parent entity.")
	iottwinmaker_updateEntityCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the entity.")
	iottwinmaker_updateEntityCmd.MarkFlagRequired("entity-id")
	iottwinmaker_updateEntityCmd.MarkFlagRequired("workspace-id")
	iottwinmakerCmd.AddCommand(iottwinmaker_updateEntityCmd)
}
