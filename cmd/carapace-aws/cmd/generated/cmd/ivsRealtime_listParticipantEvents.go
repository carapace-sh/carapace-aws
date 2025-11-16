package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_listParticipantEventsCmd = &cobra.Command{
	Use:   "list-participant-events",
	Short: "Lists events for a specified participant that occurred during a specified stage session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_listParticipantEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_listParticipantEventsCmd).Standalone()

		ivsRealtime_listParticipantEventsCmd.Flags().String("max-results", "", "Maximum number of results to return.")
		ivsRealtime_listParticipantEventsCmd.Flags().String("next-token", "", "The first participant event to retrieve.")
		ivsRealtime_listParticipantEventsCmd.Flags().String("participant-id", "", "Unique identifier for this participant.")
		ivsRealtime_listParticipantEventsCmd.Flags().String("session-id", "", "ID of a session within the stage.")
		ivsRealtime_listParticipantEventsCmd.Flags().String("stage-arn", "", "Stage ARN.")
		ivsRealtime_listParticipantEventsCmd.MarkFlagRequired("participant-id")
		ivsRealtime_listParticipantEventsCmd.MarkFlagRequired("session-id")
		ivsRealtime_listParticipantEventsCmd.MarkFlagRequired("stage-arn")
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_listParticipantEventsCmd)
}
