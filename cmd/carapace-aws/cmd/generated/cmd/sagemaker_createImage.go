package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createImageCmd = &cobra.Command{
	Use:   "create-image",
	Short: "Creates a custom SageMaker AI image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createImageCmd).Standalone()

	sagemaker_createImageCmd.Flags().String("description", "", "The description of the image.")
	sagemaker_createImageCmd.Flags().String("display-name", "", "The display name of the image.")
	sagemaker_createImageCmd.Flags().String("image-name", "", "The name of the image.")
	sagemaker_createImageCmd.Flags().String("role-arn", "", "The ARN of an IAM role that enables Amazon SageMaker AI to perform tasks on your behalf.")
	sagemaker_createImageCmd.Flags().String("tags", "", "A list of tags to apply to the image.")
	sagemaker_createImageCmd.MarkFlagRequired("image-name")
	sagemaker_createImageCmd.MarkFlagRequired("role-arn")
	sagemakerCmd.AddCommand(sagemaker_createImageCmd)
}
