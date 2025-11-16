package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_deleteContainerRecipeCmd = &cobra.Command{
	Use:   "delete-container-recipe",
	Short: "Deletes a container recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_deleteContainerRecipeCmd).Standalone()

	imagebuilder_deleteContainerRecipeCmd.Flags().String("container-recipe-arn", "", "The Amazon Resource Name (ARN) of the container recipe to delete.")
	imagebuilder_deleteContainerRecipeCmd.MarkFlagRequired("container-recipe-arn")
	imagebuilderCmd.AddCommand(imagebuilder_deleteContainerRecipeCmd)
}
