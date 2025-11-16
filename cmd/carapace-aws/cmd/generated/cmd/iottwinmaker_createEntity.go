package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iottwinmaker_createEntityCmd = &cobra.Command{
	Use:   "create-entity",
	Short: "Creates an entity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iottwinmaker_createEntityCmd).Standalone()

	iottwinmaker_createEntityCmd.Flags().String("components", "", "An object that maps strings to the components in the entity.")
	iottwinmaker_createEntityCmd.Flags().String("composite-components", "", "This is an object that maps strings to `compositeComponent` updates in the request.")
	iottwinmaker_createEntityCmd.Flags().String("description", "", "The description of the entity.")
	iottwinmaker_createEntityCmd.Flags().String("entity-id", "", "The ID of the entity.")
	iottwinmaker_createEntityCmd.Flags().String("entity-name", "", "The name of the entity.")
	iottwinmaker_createEntityCmd.Flags().String("parent-entity-id", "", "The ID of the entity's parent entity.")
	iottwinmaker_createEntityCmd.Flags().String("tags", "", "Metadata that you can use to manage the entity.")
	iottwinmaker_createEntityCmd.Flags().String("workspace-id", "", "The ID of the workspace that contains the entity.")
	iottwinmaker_createEntityCmd.MarkFlagRequired("entity-name")
	iottwinmaker_createEntityCmd.MarkFlagRequired("workspace-id")
	iottwinmakerCmd.AddCommand(iottwinmaker_createEntityCmd)
}
