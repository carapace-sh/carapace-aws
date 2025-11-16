package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_createPartnerEventSourceCmd = &cobra.Command{
	Use:   "create-partner-event-source",
	Short: "Called by an SaaS partner to create a partner event source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_createPartnerEventSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_createPartnerEventSourceCmd).Standalone()

		events_createPartnerEventSourceCmd.Flags().String("account", "", "The Amazon Web Services account ID that is permitted to create a matching partner event bus for this partner event source.")
		events_createPartnerEventSourceCmd.Flags().String("name", "", "The name of the partner event source.")
		events_createPartnerEventSourceCmd.MarkFlagRequired("account")
		events_createPartnerEventSourceCmd.MarkFlagRequired("name")
	})
	eventsCmd.AddCommand(events_createPartnerEventSourceCmd)
}
