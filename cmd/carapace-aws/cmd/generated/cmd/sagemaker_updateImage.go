package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateImageCmd = &cobra.Command{
	Use:   "update-image",
	Short: "Updates the properties of a SageMaker AI image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateImageCmd).Standalone()

		sagemaker_updateImageCmd.Flags().String("delete-properties", "", "A list of properties to delete.")
		sagemaker_updateImageCmd.Flags().String("description", "", "The new description for the image.")
		sagemaker_updateImageCmd.Flags().String("display-name", "", "The new display name for the image.")
		sagemaker_updateImageCmd.Flags().String("image-name", "", "The name of the image to update.")
		sagemaker_updateImageCmd.Flags().String("role-arn", "", "The new ARN for the IAM role that enables Amazon SageMaker AI to perform tasks on your behalf.")
		sagemaker_updateImageCmd.MarkFlagRequired("image-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateImageCmd)
}
