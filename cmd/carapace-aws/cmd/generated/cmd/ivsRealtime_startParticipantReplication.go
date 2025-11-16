package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_startParticipantReplicationCmd = &cobra.Command{
	Use:   "start-participant-replication",
	Short: "Starts replicating a publishing participant from a source stage to a destination stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_startParticipantReplicationCmd).Standalone()

	ivsRealtime_startParticipantReplicationCmd.Flags().String("attributes", "", "Application-provided attributes to set on the replicated participant in the destination stage.")
	ivsRealtime_startParticipantReplicationCmd.Flags().String("destination-stage-arn", "", "ARN of the stage to which the participant will be replicated.")
	ivsRealtime_startParticipantReplicationCmd.Flags().String("participant-id", "", "Participant ID of the publisher that will be replicated.")
	ivsRealtime_startParticipantReplicationCmd.Flags().String("reconnect-window-seconds", "", "If the participant disconnects and then reconnects within the specified interval, replication will continue to be `ACTIVE`.")
	ivsRealtime_startParticipantReplicationCmd.Flags().String("source-stage-arn", "", "ARN of the stage where the participant is publishing.")
	ivsRealtime_startParticipantReplicationCmd.MarkFlagRequired("destination-stage-arn")
	ivsRealtime_startParticipantReplicationCmd.MarkFlagRequired("participant-id")
	ivsRealtime_startParticipantReplicationCmd.MarkFlagRequired("source-stage-arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_startParticipantReplicationCmd)
}
