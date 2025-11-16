package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivsRealtime_listParticipantsCmd = &cobra.Command{
	Use:   "list-participants",
	Short: "Lists all participants in a specified stage session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivsRealtime_listParticipantsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivsRealtime_listParticipantsCmd).Standalone()

		ivsRealtime_listParticipantsCmd.Flags().String("filter-by-published", "", "Filters the response list to only show participants who published during the stage session.")
		ivsRealtime_listParticipantsCmd.Flags().String("filter-by-recording-state", "", "Filters the response list to only show participants with the specified recording state.")
		ivsRealtime_listParticipantsCmd.Flags().String("filter-by-state", "", "Filters the response list to only show participants in the specified state.")
		ivsRealtime_listParticipantsCmd.Flags().String("filter-by-user-id", "", "Filters the response list to match the specified user ID.")
		ivsRealtime_listParticipantsCmd.Flags().String("max-results", "", "Maximum number of results to return.")
		ivsRealtime_listParticipantsCmd.Flags().String("next-token", "", "The first participant to retrieve.")
		ivsRealtime_listParticipantsCmd.Flags().String("session-id", "", "ID of the session within the stage.")
		ivsRealtime_listParticipantsCmd.Flags().String("stage-arn", "", "Stage ARN.")
		ivsRealtime_listParticipantsCmd.MarkFlagRequired("session-id")
		ivsRealtime_listParticipantsCmd.MarkFlagRequired("stage-arn")
	})
	ivsRealtimeCmd.AddCommand(ivsRealtime_listParticipantsCmd)
}
