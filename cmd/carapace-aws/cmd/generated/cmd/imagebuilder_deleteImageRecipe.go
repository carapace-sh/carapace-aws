package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_deleteImageRecipeCmd = &cobra.Command{
	Use:   "delete-image-recipe",
	Short: "Deletes an image recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_deleteImageRecipeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_deleteImageRecipeCmd).Standalone()

		imagebuilder_deleteImageRecipeCmd.Flags().String("image-recipe-arn", "", "The Amazon Resource Name (ARN) of the image recipe to delete.")
		imagebuilder_deleteImageRecipeCmd.MarkFlagRequired("image-recipe-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_deleteImageRecipeCmd)
}
