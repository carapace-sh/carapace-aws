package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_createPipelineCmd = &cobra.Command{
	Use:   "create-pipeline",
	Short: "The CreatePipeline operation creates a pipeline with settings that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_createPipelineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elastictranscoder_createPipelineCmd).Standalone()

		elastictranscoder_createPipelineCmd.Flags().String("aws-kms-key-arn", "", "The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.")
		elastictranscoder_createPipelineCmd.Flags().String("content-config", "", "The optional `ContentConfig` object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists: which bucket to use, which users you want to have access to the files, the type of access you want users to have, and the storage class that you want to assign to the files.")
		elastictranscoder_createPipelineCmd.Flags().String("input-bucket", "", "The Amazon S3 bucket in which you saved the media files that you want to transcode.")
		elastictranscoder_createPipelineCmd.Flags().String("name", "", "The name of the pipeline.")
		elastictranscoder_createPipelineCmd.Flags().String("notifications", "", "The Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status.")
		elastictranscoder_createPipelineCmd.Flags().String("output-bucket", "", "The Amazon S3 bucket in which you want Elastic Transcoder to save the transcoded files.")
		elastictranscoder_createPipelineCmd.Flags().String("role", "", "The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to create the pipeline.")
		elastictranscoder_createPipelineCmd.Flags().String("thumbnail-config", "", "The `ThumbnailConfig` object specifies several values, including the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files, which users you want to have access to the files, the type of access you want users to have, and the storage class that you want to assign to the files.")
		elastictranscoder_createPipelineCmd.MarkFlagRequired("input-bucket")
		elastictranscoder_createPipelineCmd.MarkFlagRequired("name")
		elastictranscoder_createPipelineCmd.MarkFlagRequired("role")
	})
	elastictranscoderCmd.AddCommand(elastictranscoder_createPipelineCmd)
}
