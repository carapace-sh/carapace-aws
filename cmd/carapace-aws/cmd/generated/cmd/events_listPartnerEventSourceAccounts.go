package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listPartnerEventSourceAccountsCmd = &cobra.Command{
	Use:   "list-partner-event-source-accounts",
	Short: "An SaaS partner can use this operation to display the Amazon Web Services account ID that a particular partner event source name is associated with.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listPartnerEventSourceAccountsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_listPartnerEventSourceAccountsCmd).Standalone()

		events_listPartnerEventSourceAccountsCmd.Flags().String("event-source-name", "", "The name of the partner event source to display account information about.")
		events_listPartnerEventSourceAccountsCmd.Flags().String("limit", "", "Specifying this limits the number of results returned by this operation.")
		events_listPartnerEventSourceAccountsCmd.Flags().String("next-token", "", "The token returned by a previous call, which you can use to retrieve the next set of results.")
		events_listPartnerEventSourceAccountsCmd.MarkFlagRequired("event-source-name")
	})
	eventsCmd.AddCommand(events_listPartnerEventSourceAccountsCmd)
}
