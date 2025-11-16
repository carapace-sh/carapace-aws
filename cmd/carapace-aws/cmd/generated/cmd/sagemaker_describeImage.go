package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeImageCmd = &cobra.Command{
	Use:   "describe-image",
	Short: "Describes a SageMaker AI image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeImageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeImageCmd).Standalone()

		sagemaker_describeImageCmd.Flags().String("image-name", "", "The name of the image to describe.")
		sagemaker_describeImageCmd.MarkFlagRequired("image-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeImageCmd)
}
