package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeImageVersionCmd = &cobra.Command{
	Use:   "describe-image-version",
	Short: "Describes a version of a SageMaker AI image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeImageVersionCmd).Standalone()

	sagemaker_describeImageVersionCmd.Flags().String("alias", "", "The alias of the image version.")
	sagemaker_describeImageVersionCmd.Flags().String("image-name", "", "The name of the image.")
	sagemaker_describeImageVersionCmd.Flags().String("version", "", "The version of the image.")
	sagemaker_describeImageVersionCmd.MarkFlagRequired("image-name")
	sagemakerCmd.AddCommand(sagemaker_describeImageVersionCmd)
}
