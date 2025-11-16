package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_createImageRecipeCmd = &cobra.Command{
	Use:   "create-image-recipe",
	Short: "Creates a new image recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_createImageRecipeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_createImageRecipeCmd).Standalone()

		imagebuilder_createImageRecipeCmd.Flags().String("additional-instance-configuration", "", "Specify additional settings and launch scripts for your build instances.")
		imagebuilder_createImageRecipeCmd.Flags().String("ami-tags", "", "Tags that are applied to the AMI that Image Builder creates during the Build phase prior to image distribution.")
		imagebuilder_createImageRecipeCmd.Flags().String("block-device-mappings", "", "The block device mappings of the image recipe.")
		imagebuilder_createImageRecipeCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_createImageRecipeCmd.Flags().String("components", "", "The components included in the image recipe.")
		imagebuilder_createImageRecipeCmd.Flags().String("description", "", "The description of the image recipe.")
		imagebuilder_createImageRecipeCmd.Flags().String("name", "", "The name of the image recipe.")
		imagebuilder_createImageRecipeCmd.Flags().String("parent-image", "", "The base image for customizations specified in the image recipe.")
		imagebuilder_createImageRecipeCmd.Flags().String("semantic-version", "", "The semantic version of the image recipe.")
		imagebuilder_createImageRecipeCmd.Flags().String("tags", "", "The tags of the image recipe.")
		imagebuilder_createImageRecipeCmd.Flags().String("working-directory", "", "The working directory used during build and test workflows.")
		imagebuilder_createImageRecipeCmd.MarkFlagRequired("client-token")
		imagebuilder_createImageRecipeCmd.MarkFlagRequired("name")
		imagebuilder_createImageRecipeCmd.MarkFlagRequired("parent-image")
		imagebuilder_createImageRecipeCmd.MarkFlagRequired("semantic-version")
	})
	imagebuilderCmd.AddCommand(imagebuilder_createImageRecipeCmd)
}
