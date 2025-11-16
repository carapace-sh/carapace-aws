package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchQueuesCmd = &cobra.Command{
	Use:   "search-queues",
	Short: "Searches queues in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchQueuesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_searchQueuesCmd).Standalone()

		connect_searchQueuesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_searchQueuesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_searchQueuesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_searchQueuesCmd.Flags().String("search-criteria", "", "The search criteria to be used to return queues.")
		connect_searchQueuesCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
		connect_searchQueuesCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_searchQueuesCmd)
}
