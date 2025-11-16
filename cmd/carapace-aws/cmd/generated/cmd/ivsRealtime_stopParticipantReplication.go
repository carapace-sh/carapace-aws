package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_stopParticipantReplicationCmd = &cobra.Command{
	Use:   "stop-participant-replication",
	Short: "Stops a replicated participant session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_stopParticipantReplicationCmd).Standalone()

	ivsRealtime_stopParticipantReplicationCmd.Flags().String("destination-stage-arn", "", "ARN of the stage where the participant has been replicated.")
	ivsRealtime_stopParticipantReplicationCmd.Flags().String("participant-id", "", "Participant ID of the publisher that has been replicated.")
	ivsRealtime_stopParticipantReplicationCmd.Flags().String("source-stage-arn", "", "ARN of the stage where the participant is publishing.")
	ivsRealtime_stopParticipantReplicationCmd.MarkFlagRequired("destination-stage-arn")
	ivsRealtime_stopParticipantReplicationCmd.MarkFlagRequired("participant-id")
	ivsRealtime_stopParticipantReplicationCmd.MarkFlagRequired("source-stage-arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_stopParticipantReplicationCmd)
}
