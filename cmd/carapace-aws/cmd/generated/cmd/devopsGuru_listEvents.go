package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_listEventsCmd = &cobra.Command{
	Use:   "list-events",
	Short: "Returns a list of the events emitted by the resources that are evaluated by DevOps Guru.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_listEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_listEventsCmd).Standalone()

		devopsGuru_listEventsCmd.Flags().String("account-id", "", "The ID of the Amazon Web Services account.")
		devopsGuru_listEventsCmd.Flags().String("filters", "", "A `ListEventsFilters` object used to specify which events to return.")
		devopsGuru_listEventsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		devopsGuru_listEventsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
		devopsGuru_listEventsCmd.MarkFlagRequired("filters")
	})
	devopsGuruCmd.AddCommand(devopsGuru_listEventsCmd)
}
