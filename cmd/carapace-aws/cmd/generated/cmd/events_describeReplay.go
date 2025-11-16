package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_describeReplayCmd = &cobra.Command{
	Use:   "describe-replay",
	Short: "Retrieves details about a replay.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_describeReplayCmd).Standalone()

	events_describeReplayCmd.Flags().String("replay-name", "", "The name of the replay to retrieve.")
	events_describeReplayCmd.MarkFlagRequired("replay-name")
	eventsCmd.AddCommand(events_describeReplayCmd)
}
