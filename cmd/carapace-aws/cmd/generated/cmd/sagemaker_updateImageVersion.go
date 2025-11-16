package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateImageVersionCmd = &cobra.Command{
	Use:   "update-image-version",
	Short: "Updates the properties of a SageMaker AI image version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateImageVersionCmd).Standalone()

	sagemaker_updateImageVersionCmd.Flags().String("alias", "", "The alias of the image version.")
	sagemaker_updateImageVersionCmd.Flags().String("aliases-to-add", "", "A list of aliases to add.")
	sagemaker_updateImageVersionCmd.Flags().String("aliases-to-delete", "", "A list of aliases to delete.")
	sagemaker_updateImageVersionCmd.Flags().String("horovod", "", "Indicates Horovod compatibility.")
	sagemaker_updateImageVersionCmd.Flags().String("image-name", "", "The name of the image.")
	sagemaker_updateImageVersionCmd.Flags().String("job-type", "", "Indicates SageMaker AI job type compatibility.")
	sagemaker_updateImageVersionCmd.Flags().String("mlframework", "", "The machine learning framework vended in the image version.")
	sagemaker_updateImageVersionCmd.Flags().String("processor", "", "Indicates CPU or GPU compatibility.")
	sagemaker_updateImageVersionCmd.Flags().String("programming-lang", "", "The supported programming language and its version.")
	sagemaker_updateImageVersionCmd.Flags().String("release-notes", "", "The maintainer description of the image version.")
	sagemaker_updateImageVersionCmd.Flags().String("vendor-guidance", "", "The availability of the image version specified by the maintainer.")
	sagemaker_updateImageVersionCmd.Flags().String("version", "", "The version of the image.")
	sagemaker_updateImageVersionCmd.MarkFlagRequired("image-name")
	sagemakerCmd.AddCommand(sagemaker_updateImageVersionCmd)
}
