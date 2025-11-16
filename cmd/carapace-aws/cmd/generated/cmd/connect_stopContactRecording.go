package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_stopContactRecordingCmd = &cobra.Command{
	Use:   "stop-contact-recording",
	Short: "Stops recording a call when a contact is being recorded.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_stopContactRecordingCmd).Standalone()

	connect_stopContactRecordingCmd.Flags().String("contact-id", "", "The identifier of the contact.")
	connect_stopContactRecordingCmd.Flags().String("contact-recording-type", "", "The type of recording being operated on.")
	connect_stopContactRecordingCmd.Flags().String("initial-contact-id", "", "The identifier of the contact.")
	connect_stopContactRecordingCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_stopContactRecordingCmd.MarkFlagRequired("contact-id")
	connect_stopContactRecordingCmd.MarkFlagRequired("initial-contact-id")
	connect_stopContactRecordingCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_stopContactRecordingCmd)
}
