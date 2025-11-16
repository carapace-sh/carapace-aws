package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rekognition_createStreamProcessorCmd = &cobra.Command{
	Use:   "create-stream-processor",
	Short: "Creates an Amazon Rekognition stream processor that you can use to detect and recognize faces or to detect labels in a streaming video.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rekognition_createStreamProcessorCmd).Standalone()

	rekognition_createStreamProcessorCmd.Flags().String("data-sharing-preference", "", "Shows whether you are sharing data with Rekognition to improve model performance.")
	rekognition_createStreamProcessorCmd.Flags().String("input", "", "Kinesis video stream stream that provides the source streaming video.")
	rekognition_createStreamProcessorCmd.Flags().String("kms-key-id", "", "The identifier for your AWS Key Management Service key (AWS KMS key).")
	rekognition_createStreamProcessorCmd.Flags().String("name", "", "An identifier you assign to the stream processor.")
	rekognition_createStreamProcessorCmd.Flags().String("notification-channel", "", "")
	rekognition_createStreamProcessorCmd.Flags().String("output", "", "Kinesis data stream stream or Amazon S3 bucket location to which Amazon Rekognition Video puts the analysis results.")
	rekognition_createStreamProcessorCmd.Flags().String("regions-of-interest", "", "Specifies locations in the frames where Amazon Rekognition checks for objects or people.")
	rekognition_createStreamProcessorCmd.Flags().String("role-arn", "", "The Amazon Resource Number (ARN) of the IAM role that allows access to the stream processor.")
	rekognition_createStreamProcessorCmd.Flags().String("settings", "", "Input parameters used in a streaming video analyzed by a stream processor.")
	rekognition_createStreamProcessorCmd.Flags().String("tags", "", "A set of tags (key-value pairs) that you want to attach to the stream processor.")
	rekognition_createStreamProcessorCmd.MarkFlagRequired("input")
	rekognition_createStreamProcessorCmd.MarkFlagRequired("name")
	rekognition_createStreamProcessorCmd.MarkFlagRequired("output")
	rekognition_createStreamProcessorCmd.MarkFlagRequired("role-arn")
	rekognition_createStreamProcessorCmd.MarkFlagRequired("settings")
	rekognitionCmd.AddCommand(rekognition_createStreamProcessorCmd)
}
