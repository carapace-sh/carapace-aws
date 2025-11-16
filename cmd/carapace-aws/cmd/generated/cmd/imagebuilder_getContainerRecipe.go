package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getContainerRecipeCmd = &cobra.Command{
	Use:   "get-container-recipe",
	Short: "Retrieves a container recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getContainerRecipeCmd).Standalone()

	imagebuilder_getContainerRecipeCmd.Flags().String("container-recipe-arn", "", "The Amazon Resource Name (ARN) of the container recipe to retrieve.")
	imagebuilder_getContainerRecipeCmd.MarkFlagRequired("container-recipe-arn")
	imagebuilderCmd.AddCommand(imagebuilder_getContainerRecipeCmd)
}
