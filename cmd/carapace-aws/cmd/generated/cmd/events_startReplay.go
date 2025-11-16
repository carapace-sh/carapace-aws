package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_startReplayCmd = &cobra.Command{
	Use:   "start-replay",
	Short: "Starts the specified replay.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_startReplayCmd).Standalone()

	events_startReplayCmd.Flags().String("description", "", "A description for the replay to start.")
	events_startReplayCmd.Flags().String("destination", "", "A `ReplayDestination` object that includes details about the destination for the replay.")
	events_startReplayCmd.Flags().String("event-end-time", "", "A time stamp for the time to stop replaying events.")
	events_startReplayCmd.Flags().String("event-source-arn", "", "The ARN of the archive to replay events from.")
	events_startReplayCmd.Flags().String("event-start-time", "", "A time stamp for the time to start replaying events.")
	events_startReplayCmd.Flags().String("replay-name", "", "The name of the replay to start.")
	events_startReplayCmd.MarkFlagRequired("destination")
	events_startReplayCmd.MarkFlagRequired("event-end-time")
	events_startReplayCmd.MarkFlagRequired("event-source-arn")
	events_startReplayCmd.MarkFlagRequired("event-start-time")
	events_startReplayCmd.MarkFlagRequired("replay-name")
	eventsCmd.AddCommand(events_startReplayCmd)
}
