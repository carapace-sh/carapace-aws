package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createPartnerAppPresignedUrlCmd = &cobra.Command{
	Use:   "create-partner-app-presigned-url",
	Short: "Creates a presigned URL to access an Amazon SageMaker Partner AI App.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createPartnerAppPresignedUrlCmd).Standalone()

	sagemaker_createPartnerAppPresignedUrlCmd.Flags().String("arn", "", "The ARN of the SageMaker Partner AI App to create the presigned URL for.")
	sagemaker_createPartnerAppPresignedUrlCmd.Flags().String("expires-in-seconds", "", "The time that will pass before the presigned URL expires.")
	sagemaker_createPartnerAppPresignedUrlCmd.Flags().String("session-expiration-duration-in-seconds", "", "Indicates how long the Amazon SageMaker Partner AI App session can be accessed for after logging in.")
	sagemaker_createPartnerAppPresignedUrlCmd.MarkFlagRequired("arn")
	sagemakerCmd.AddCommand(sagemaker_createPartnerAppPresignedUrlCmd)
}
