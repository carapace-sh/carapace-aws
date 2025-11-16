package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchAgentStatusesCmd = &cobra.Command{
	Use:   "search-agent-statuses",
	Short: "Searches AgentStatuses in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchAgentStatusesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_searchAgentStatusesCmd).Standalone()

		connect_searchAgentStatusesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_searchAgentStatusesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_searchAgentStatusesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_searchAgentStatusesCmd.Flags().String("search-criteria", "", "The search criteria to be used to return agent statuses.")
		connect_searchAgentStatusesCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
		connect_searchAgentStatusesCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_searchAgentStatusesCmd)
}
