package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_createContainerRecipeCmd = &cobra.Command{
	Use:   "create-container-recipe",
	Short: "Creates a new container recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_createContainerRecipeCmd).Standalone()

	imagebuilder_createContainerRecipeCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
	imagebuilder_createContainerRecipeCmd.Flags().String("components", "", "The components included in the container recipe.")
	imagebuilder_createContainerRecipeCmd.Flags().String("container-type", "", "The type of container to create.")
	imagebuilder_createContainerRecipeCmd.Flags().String("description", "", "The description of the container recipe.")
	imagebuilder_createContainerRecipeCmd.Flags().String("dockerfile-template-data", "", "The Dockerfile template used to build your image as an inline data blob.")
	imagebuilder_createContainerRecipeCmd.Flags().String("dockerfile-template-uri", "", "The Amazon S3 URI for the Dockerfile that will be used to build your container image.")
	imagebuilder_createContainerRecipeCmd.Flags().String("image-os-version-override", "", "Specifies the operating system version for the base image.")
	imagebuilder_createContainerRecipeCmd.Flags().String("instance-configuration", "", "A group of options that can be used to configure an instance for building and testing container images.")
	imagebuilder_createContainerRecipeCmd.Flags().String("kms-key-id", "", "The Amazon Resource Name (ARN) that uniquely identifies which KMS key is used to encrypt the Dockerfile template.")
	imagebuilder_createContainerRecipeCmd.Flags().String("name", "", "The name of the container recipe.")
	imagebuilder_createContainerRecipeCmd.Flags().String("parent-image", "", "The base image for the container recipe.")
	imagebuilder_createContainerRecipeCmd.Flags().String("platform-override", "", "Specifies the operating system platform when you use a custom base image.")
	imagebuilder_createContainerRecipeCmd.Flags().String("semantic-version", "", "The semantic version of the container recipe.")
	imagebuilder_createContainerRecipeCmd.Flags().String("tags", "", "Tags that are attached to the container recipe.")
	imagebuilder_createContainerRecipeCmd.Flags().String("target-repository", "", "The destination repository for the container image.")
	imagebuilder_createContainerRecipeCmd.Flags().String("working-directory", "", "The working directory for use during build and test workflows.")
	imagebuilder_createContainerRecipeCmd.MarkFlagRequired("client-token")
	imagebuilder_createContainerRecipeCmd.MarkFlagRequired("container-type")
	imagebuilder_createContainerRecipeCmd.MarkFlagRequired("name")
	imagebuilder_createContainerRecipeCmd.MarkFlagRequired("parent-image")
	imagebuilder_createContainerRecipeCmd.MarkFlagRequired("semantic-version")
	imagebuilder_createContainerRecipeCmd.MarkFlagRequired("target-repository")
	imagebuilderCmd.AddCommand(imagebuilder_createContainerRecipeCmd)
}
