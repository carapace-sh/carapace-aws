package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_deletePartnerEventSourceCmd = &cobra.Command{
	Use:   "delete-partner-event-source",
	Short: "This operation is used by SaaS partners to delete a partner event source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_deletePartnerEventSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_deletePartnerEventSourceCmd).Standalone()

		events_deletePartnerEventSourceCmd.Flags().String("account", "", "The Amazon Web Services account ID of the Amazon Web Services customer that the event source was created for.")
		events_deletePartnerEventSourceCmd.Flags().String("name", "", "The name of the event source to delete.")
		events_deletePartnerEventSourceCmd.MarkFlagRequired("account")
		events_deletePartnerEventSourceCmd.MarkFlagRequired("name")
	})
	eventsCmd.AddCommand(events_deletePartnerEventSourceCmd)
}
