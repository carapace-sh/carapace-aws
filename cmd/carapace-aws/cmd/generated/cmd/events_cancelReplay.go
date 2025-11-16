package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_cancelReplayCmd = &cobra.Command{
	Use:   "cancel-replay",
	Short: "Cancels the specified replay.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_cancelReplayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_cancelReplayCmd).Standalone()

		events_cancelReplayCmd.Flags().String("replay-name", "", "The name of the replay to cancel.")
		events_cancelReplayCmd.MarkFlagRequired("replay-name")
	})
	eventsCmd.AddCommand(events_cancelReplayCmd)
}
