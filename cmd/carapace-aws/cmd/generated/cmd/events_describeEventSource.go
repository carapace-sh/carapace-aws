package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_describeEventSourceCmd = &cobra.Command{
	Use:   "describe-event-source",
	Short: "This operation lists details about a partner event source that is shared with your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_describeEventSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_describeEventSourceCmd).Standalone()

		events_describeEventSourceCmd.Flags().String("name", "", "The name of the partner event source to display the details of.")
		events_describeEventSourceCmd.MarkFlagRequired("name")
	})
	eventsCmd.AddCommand(events_describeEventSourceCmd)
}
