package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createHubContentPresignedUrlsCmd = &cobra.Command{
	Use:   "create-hub-content-presigned-urls",
	Short: "Creates presigned URLs for accessing hub content artifacts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createHubContentPresignedUrlsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createHubContentPresignedUrlsCmd).Standalone()

		sagemaker_createHubContentPresignedUrlsCmd.Flags().String("access-config", "", "Configuration settings for accessing the hub content, including end-user license agreement acceptance for gated models and expected S3 URL validation.")
		sagemaker_createHubContentPresignedUrlsCmd.Flags().String("hub-content-name", "", "The name of the hub content for which to generate presigned URLs.")
		sagemaker_createHubContentPresignedUrlsCmd.Flags().String("hub-content-type", "", "The type of hub content to access.")
		sagemaker_createHubContentPresignedUrlsCmd.Flags().String("hub-content-version", "", "The version of the hub content.")
		sagemaker_createHubContentPresignedUrlsCmd.Flags().String("hub-name", "", "The name or Amazon Resource Name (ARN) of the hub that contains the content.")
		sagemaker_createHubContentPresignedUrlsCmd.Flags().String("max-results", "", "The maximum number of presigned URLs to return in the response.")
		sagemaker_createHubContentPresignedUrlsCmd.Flags().String("next-token", "", "A token for pagination.")
		sagemaker_createHubContentPresignedUrlsCmd.MarkFlagRequired("hub-content-name")
		sagemaker_createHubContentPresignedUrlsCmd.MarkFlagRequired("hub-content-type")
		sagemaker_createHubContentPresignedUrlsCmd.MarkFlagRequired("hub-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createHubContentPresignedUrlsCmd)
}
