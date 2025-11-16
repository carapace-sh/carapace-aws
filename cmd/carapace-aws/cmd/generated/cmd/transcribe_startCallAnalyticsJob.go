package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_startCallAnalyticsJobCmd = &cobra.Command{
	Use:   "start-call-analytics-job",
	Short: "Transcribes the audio from a customer service call and applies any additional Request Parameters you choose to include in your request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_startCallAnalyticsJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_startCallAnalyticsJobCmd).Standalone()

		transcribe_startCallAnalyticsJobCmd.Flags().String("call-analytics-job-name", "", "A unique name, chosen by you, for your Call Analytics job.")
		transcribe_startCallAnalyticsJobCmd.Flags().String("channel-definitions", "", "Makes it possible to specify which speaker is on which channel.")
		transcribe_startCallAnalyticsJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that has permissions to access the Amazon S3 bucket that contains your input files.")
		transcribe_startCallAnalyticsJobCmd.Flags().String("media", "", "Describes the Amazon S3 location of the media file you want to use in your Call Analytics request.")
		transcribe_startCallAnalyticsJobCmd.Flags().String("output-encryption-kmskey-id", "", "The Amazon Resource Name (ARN) of a KMS key that you want to use to encrypt your Call Analytics output.")
		transcribe_startCallAnalyticsJobCmd.Flags().String("output-location", "", "The Amazon S3 location where you want your Call Analytics transcription output stored.")
		transcribe_startCallAnalyticsJobCmd.Flags().String("settings", "", "Specify additional optional settings in your request, including content redaction; allows you to apply custom language models, vocabulary filters, and custom vocabularies to your Call Analytics job.")
		transcribe_startCallAnalyticsJobCmd.Flags().String("tags", "", "Adds one or more custom tags, each in the form of a key:value pair, to a new call analytics job at the time you start this new job.")
		transcribe_startCallAnalyticsJobCmd.MarkFlagRequired("call-analytics-job-name")
		transcribe_startCallAnalyticsJobCmd.MarkFlagRequired("media")
	})
	transcribeCmd.AddCommand(transcribe_startCallAnalyticsJobCmd)
}
