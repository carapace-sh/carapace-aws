package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_describePartnerEventSourceCmd = &cobra.Command{
	Use:   "describe-partner-event-source",
	Short: "An SaaS partner can use this operation to list details about a partner event source that they have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_describePartnerEventSourceCmd).Standalone()

	events_describePartnerEventSourceCmd.Flags().String("name", "", "The name of the event source to display.")
	events_describePartnerEventSourceCmd.MarkFlagRequired("name")
	eventsCmd.AddCommand(events_describePartnerEventSourceCmd)
}
