package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_listParticipantReplicasCmd = &cobra.Command{
	Use:   "list-participant-replicas",
	Short: "Lists all the replicas for a participant from a source stage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_listParticipantReplicasCmd).Standalone()

	ivsRealtime_listParticipantReplicasCmd.Flags().String("max-results", "", "Maximum number of results to return.")
	ivsRealtime_listParticipantReplicasCmd.Flags().String("next-token", "", "The first participant to retrieve.")
	ivsRealtime_listParticipantReplicasCmd.Flags().String("participant-id", "", "Participant ID of the publisher that has been replicated.")
	ivsRealtime_listParticipantReplicasCmd.Flags().String("source-stage-arn", "", "ARN of the stage where the participant is publishing.")
	ivsRealtime_listParticipantReplicasCmd.MarkFlagRequired("participant-id")
	ivsRealtime_listParticipantReplicasCmd.MarkFlagRequired("source-stage-arn")
	ivsRealtimeCmd.AddCommand(ivsRealtime_listParticipantReplicasCmd)
}
