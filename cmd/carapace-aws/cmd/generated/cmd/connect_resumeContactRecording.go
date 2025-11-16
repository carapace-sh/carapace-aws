package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_resumeContactRecordingCmd = &cobra.Command{
	Use:   "resume-contact-recording",
	Short: "When a contact is being recorded, and the recording has been suspended using SuspendContactRecording, this API resumes recording whatever recording is selected in the flow configuration: call, screen, or both.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_resumeContactRecordingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_resumeContactRecordingCmd).Standalone()

		connect_resumeContactRecordingCmd.Flags().String("contact-id", "", "The identifier of the contact.")
		connect_resumeContactRecordingCmd.Flags().String("contact-recording-type", "", "The type of recording being operated on.")
		connect_resumeContactRecordingCmd.Flags().String("initial-contact-id", "", "The identifier of the contact.")
		connect_resumeContactRecordingCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_resumeContactRecordingCmd.MarkFlagRequired("contact-id")
		connect_resumeContactRecordingCmd.MarkFlagRequired("initial-contact-id")
		connect_resumeContactRecordingCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_resumeContactRecordingCmd)
}
