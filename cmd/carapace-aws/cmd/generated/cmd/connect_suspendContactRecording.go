package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_suspendContactRecordingCmd = &cobra.Command{
	Use:   "suspend-contact-recording",
	Short: "When a contact is being recorded, this API suspends recording whatever is selected in the flow configuration: call (IVR or agent), screen, or both.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_suspendContactRecordingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_suspendContactRecordingCmd).Standalone()

		connect_suspendContactRecordingCmd.Flags().String("contact-id", "", "The identifier of the contact.")
		connect_suspendContactRecordingCmd.Flags().String("contact-recording-type", "", "The type of recording being operated on.")
		connect_suspendContactRecordingCmd.Flags().String("initial-contact-id", "", "The identifier of the contact.")
		connect_suspendContactRecordingCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_suspendContactRecordingCmd.MarkFlagRequired("contact-id")
		connect_suspendContactRecordingCmd.MarkFlagRequired("initial-contact-id")
		connect_suspendContactRecordingCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_suspendContactRecordingCmd)
}
