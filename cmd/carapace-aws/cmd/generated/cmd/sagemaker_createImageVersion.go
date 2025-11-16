package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createImageVersionCmd = &cobra.Command{
	Use:   "create-image-version",
	Short: "Creates a version of the SageMaker AI image specified by `ImageName`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createImageVersionCmd).Standalone()

	sagemaker_createImageVersionCmd.Flags().String("aliases", "", "A list of aliases created with the image version.")
	sagemaker_createImageVersionCmd.Flags().String("base-image", "", "The registry path of the container image to use as the starting point for this version.")
	sagemaker_createImageVersionCmd.Flags().String("client-token", "", "A unique ID.")
	sagemaker_createImageVersionCmd.Flags().String("horovod", "", "Indicates Horovod compatibility.")
	sagemaker_createImageVersionCmd.Flags().String("image-name", "", "The `ImageName` of the `Image` to create a version of.")
	sagemaker_createImageVersionCmd.Flags().String("job-type", "", "Indicates SageMaker AI job type compatibility.")
	sagemaker_createImageVersionCmd.Flags().String("mlframework", "", "The machine learning framework vended in the image version.")
	sagemaker_createImageVersionCmd.Flags().String("processor", "", "Indicates CPU or GPU compatibility.")
	sagemaker_createImageVersionCmd.Flags().String("programming-lang", "", "The supported programming language and its version.")
	sagemaker_createImageVersionCmd.Flags().String("release-notes", "", "The maintainer description of the image version.")
	sagemaker_createImageVersionCmd.Flags().String("vendor-guidance", "", "The stability of the image version, specified by the maintainer.")
	sagemaker_createImageVersionCmd.MarkFlagRequired("base-image")
	sagemaker_createImageVersionCmd.MarkFlagRequired("client-token")
	sagemaker_createImageVersionCmd.MarkFlagRequired("image-name")
	sagemakerCmd.AddCommand(sagemaker_createImageVersionCmd)
}
