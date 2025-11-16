package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_disconnectParticipantCmd = &cobra.Command{
	Use:   "disconnect-participant",
	Short: "Disconnects a specified participant from a specified stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_disconnectParticipantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_disconnectParticipantCmd).Standalone()

		ivsRealtime_disconnectParticipantCmd.Flags().String("participant-id", "", "Identifier of the participant to be disconnected.")
		ivsRealtime_disconnectParticipantCmd.Flags().String("reason", "", "Description of why this participant is being disconnected.")
		ivsRealtime_disconnectParticipantCmd.Flags().String("stage-arn", "", "ARN of the stage to which the participant is attached.")
		ivsRealtime_disconnectParticipantCmd.MarkFlagRequired("participant-id")
		ivsRealtime_disconnectParticipantCmd.MarkFlagRequired("stage-arn")
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_disconnectParticipantCmd)
}
