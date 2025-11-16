package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteImageCmd = &cobra.Command{
	Use:   "delete-image",
	Short: "Deletes a SageMaker AI image and all versions of the image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteImageCmd).Standalone()

	sagemaker_deleteImageCmd.Flags().String("image-name", "", "The name of the image to delete.")
	sagemaker_deleteImageCmd.MarkFlagRequired("image-name")
	sagemakerCmd.AddCommand(sagemaker_deleteImageCmd)
}
