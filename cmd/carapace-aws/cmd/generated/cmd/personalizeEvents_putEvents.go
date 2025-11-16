package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalizeEvents_putEventsCmd = &cobra.Command{
	Use:   "put-events",
	Short: "Records item interaction event data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalizeEvents_putEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalizeEvents_putEventsCmd).Standalone()

		personalizeEvents_putEventsCmd.Flags().String("event-list", "", "A list of event data from the session.")
		personalizeEvents_putEventsCmd.Flags().String("session-id", "", "The session ID associated with the user's visit.")
		personalizeEvents_putEventsCmd.Flags().String("tracking-id", "", "The tracking ID for the event.")
		personalizeEvents_putEventsCmd.Flags().String("user-id", "", "The user associated with the event.")
		personalizeEvents_putEventsCmd.MarkFlagRequired("event-list")
		personalizeEvents_putEventsCmd.MarkFlagRequired("session-id")
		personalizeEvents_putEventsCmd.MarkFlagRequired("tracking-id")
	})
	personalizeEventsCmd.AddCommand(personalizeEvents_putEventsCmd)
}
