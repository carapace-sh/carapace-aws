package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_getParticipantCmd = &cobra.Command{
	Use:   "get-participant",
	Short: "Gets information about the specified participant token.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_getParticipantCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_getParticipantCmd).Standalone()

		ivsRealtime_getParticipantCmd.Flags().String("participant-id", "", "Unique identifier for the participant.")
		ivsRealtime_getParticipantCmd.Flags().String("session-id", "", "ID of a session within the stage.")
		ivsRealtime_getParticipantCmd.Flags().String("stage-arn", "", "Stage ARN.")
		ivsRealtime_getParticipantCmd.MarkFlagRequired("participant-id")
		ivsRealtime_getParticipantCmd.MarkFlagRequired("session-id")
		ivsRealtime_getParticipantCmd.MarkFlagRequired("stage-arn")
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_getParticipantCmd)
}
