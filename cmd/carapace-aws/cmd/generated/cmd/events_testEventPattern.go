package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_testEventPatternCmd = &cobra.Command{
	Use:   "test-event-pattern",
	Short: "Tests whether the specified event pattern matches the provided event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_testEventPatternCmd).Standalone()

	events_testEventPatternCmd.Flags().String("event", "", "The event, in JSON format, to test against the event pattern.")
	events_testEventPatternCmd.Flags().String("event-pattern", "", "The event pattern.")
	events_testEventPatternCmd.MarkFlagRequired("event")
	events_testEventPatternCmd.MarkFlagRequired("event-pattern")
	eventsCmd.AddCommand(events_testEventPatternCmd)
}
