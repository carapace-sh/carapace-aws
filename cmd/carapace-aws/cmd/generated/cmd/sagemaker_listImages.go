package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listImagesCmd = &cobra.Command{
	Use:   "list-images",
	Short: "Lists the images in your account and their properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listImagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listImagesCmd).Standalone()

		sagemaker_listImagesCmd.Flags().String("creation-time-after", "", "A filter that returns only images created on or after the specified time.")
		sagemaker_listImagesCmd.Flags().String("creation-time-before", "", "A filter that returns only images created on or before the specified time.")
		sagemaker_listImagesCmd.Flags().String("last-modified-time-after", "", "A filter that returns only images modified on or after the specified time.")
		sagemaker_listImagesCmd.Flags().String("last-modified-time-before", "", "A filter that returns only images modified on or before the specified time.")
		sagemaker_listImagesCmd.Flags().String("max-results", "", "The maximum number of images to return in the response.")
		sagemaker_listImagesCmd.Flags().String("name-contains", "", "A filter that returns only images whose name contains the specified string.")
		sagemaker_listImagesCmd.Flags().String("next-token", "", "If the previous call to `ListImages` didn't return the full set of images, the call returns a token for getting the next set of images.")
		sagemaker_listImagesCmd.Flags().String("sort-by", "", "The property used to sort results.")
		sagemaker_listImagesCmd.Flags().String("sort-order", "", "The sort order.")
	})
	sagemakerCmd.AddCommand(sagemaker_listImagesCmd)
}
