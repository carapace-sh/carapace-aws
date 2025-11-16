package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elastictranscoder_updatePipelineCmd = &cobra.Command{
	Use:   "update-pipeline",
	Short: "Use the `UpdatePipeline` operation to update settings for a pipeline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elastictranscoder_updatePipelineCmd).Standalone()

	elastictranscoder_updatePipelineCmd.Flags().String("aws-kms-key-arn", "", "The AWS Key Management Service (AWS KMS) key that you want to use with this pipeline.")
	elastictranscoder_updatePipelineCmd.Flags().String("content-config", "", "The optional `ContentConfig` object specifies information about the Amazon S3 bucket in which you want Elastic Transcoder to save transcoded files and playlists: which bucket to use, which users you want to have access to the files, the type of access you want users to have, and the storage class that you want to assign to the files.")
	elastictranscoder_updatePipelineCmd.Flags().String("id", "", "The ID of the pipeline that you want to update.")
	elastictranscoder_updatePipelineCmd.Flags().String("input-bucket", "", "The Amazon S3 bucket in which you saved the media files that you want to transcode and the graphics that you want to use as watermarks.")
	elastictranscoder_updatePipelineCmd.Flags().String("name", "", "The name of the pipeline.")
	elastictranscoder_updatePipelineCmd.Flags().String("notifications", "", "The topic ARN for the Amazon Simple Notification Service (Amazon SNS) topic that you want to notify to report job status.")
	elastictranscoder_updatePipelineCmd.Flags().String("role", "", "The IAM Amazon Resource Name (ARN) for the role that you want Elastic Transcoder to use to transcode jobs for this pipeline.")
	elastictranscoder_updatePipelineCmd.Flags().String("thumbnail-config", "", "The `ThumbnailConfig` object specifies several values, including the Amazon S3 bucket in which you want Elastic Transcoder to save thumbnail files, which users you want to have access to the files, the type of access you want users to have, and the storage class that you want to assign to the files.")
	elastictranscoder_updatePipelineCmd.MarkFlagRequired("id")
	elastictranscoderCmd.AddCommand(elastictranscoder_updatePipelineCmd)
}
