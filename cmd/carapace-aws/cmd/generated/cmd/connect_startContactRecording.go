package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_startContactRecordingCmd = &cobra.Command{
	Use:   "start-contact-recording",
	Short: "Starts recording the contact:\n\n- If the API is called *before* the agent joins the call, recording starts when the agent joins the call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_startContactRecordingCmd).Standalone()

	connect_startContactRecordingCmd.Flags().String("contact-id", "", "The identifier of the contact.")
	connect_startContactRecordingCmd.Flags().String("initial-contact-id", "", "The identifier of the contact.")
	connect_startContactRecordingCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_startContactRecordingCmd.Flags().String("voice-recording-configuration", "", "The person being recorded.")
	connect_startContactRecordingCmd.MarkFlagRequired("contact-id")
	connect_startContactRecordingCmd.MarkFlagRequired("initial-contact-id")
	connect_startContactRecordingCmd.MarkFlagRequired("instance-id")
	connect_startContactRecordingCmd.MarkFlagRequired("voice-recording-configuration")
	connectCmd.AddCommand(connect_startContactRecordingCmd)
}
