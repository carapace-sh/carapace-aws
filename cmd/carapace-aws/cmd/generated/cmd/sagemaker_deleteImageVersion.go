package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteImageVersionCmd = &cobra.Command{
	Use:   "delete-image-version",
	Short: "Deletes a version of a SageMaker AI image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteImageVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteImageVersionCmd).Standalone()

		sagemaker_deleteImageVersionCmd.Flags().String("alias", "", "The alias of the image to delete.")
		sagemaker_deleteImageVersionCmd.Flags().String("image-name", "", "The name of the image to delete.")
		sagemaker_deleteImageVersionCmd.Flags().String("version", "", "The version to delete.")
		sagemaker_deleteImageVersionCmd.MarkFlagRequired("image-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteImageVersionCmd)
}
