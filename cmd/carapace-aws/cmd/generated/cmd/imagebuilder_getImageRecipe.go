package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getImageRecipeCmd = &cobra.Command{
	Use:   "get-image-recipe",
	Short: "Gets an image recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getImageRecipeCmd).Standalone()

	imagebuilder_getImageRecipeCmd.Flags().String("image-recipe-arn", "", "The Amazon Resource Name (ARN) of the image recipe that you want to retrieve.")
	imagebuilder_getImageRecipeCmd.MarkFlagRequired("image-recipe-arn")
	imagebuilderCmd.AddCommand(imagebuilder_getImageRecipeCmd)
}
