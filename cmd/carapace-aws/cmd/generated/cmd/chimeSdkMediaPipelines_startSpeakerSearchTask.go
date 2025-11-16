package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkMediaPipelines_startSpeakerSearchTaskCmd = &cobra.Command{
	Use:   "start-speaker-search-task",
	Short: "Starts a speaker search task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkMediaPipelines_startSpeakerSearchTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkMediaPipelines_startSpeakerSearchTaskCmd).Standalone()

		chimeSdkMediaPipelines_startSpeakerSearchTaskCmd.Flags().String("client-request-token", "", "The unique identifier for the client request.")
		chimeSdkMediaPipelines_startSpeakerSearchTaskCmd.Flags().String("identifier", "", "The unique identifier of the resource to be updated.")
		chimeSdkMediaPipelines_startSpeakerSearchTaskCmd.Flags().String("kinesis-video-stream-source-task-configuration", "", "The task configuration for the Kinesis video stream source of the media insights pipeline.")
		chimeSdkMediaPipelines_startSpeakerSearchTaskCmd.Flags().String("voice-profile-domain-arn", "", "The ARN of the voice profile domain that will store the voice profile.")
		chimeSdkMediaPipelines_startSpeakerSearchTaskCmd.MarkFlagRequired("identifier")
		chimeSdkMediaPipelines_startSpeakerSearchTaskCmd.MarkFlagRequired("voice-profile-domain-arn")
	})
	chimeSdkMediaPipelinesCmd.AddCommand(chimeSdkMediaPipelines_startSpeakerSearchTaskCmd)
}
