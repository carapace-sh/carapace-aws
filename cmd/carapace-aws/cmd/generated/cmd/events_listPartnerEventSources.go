package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var events_listPartnerEventSourcesCmd = &cobra.Command{
	Use:   "list-partner-event-sources",
	Short: "An SaaS partner can use this operation to list all the partner event source names that they have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(events_listPartnerEventSourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(events_listPartnerEventSourcesCmd).Standalone()

		events_listPartnerEventSourcesCmd.Flags().String("limit", "", "pecifying this limits the number of results returned by this operation.")
		events_listPartnerEventSourcesCmd.Flags().String("name-prefix", "", "If you specify this, the results are limited to only those partner event sources that start with the string you specify.")
		events_listPartnerEventSourcesCmd.Flags().String("next-token", "", "The token returned by a previous call, which you can use to retrieve the next set of results.")
		events_listPartnerEventSourcesCmd.MarkFlagRequired("name-prefix")
	})
	eventsCmd.AddCommand(events_listPartnerEventSourcesCmd)
}
