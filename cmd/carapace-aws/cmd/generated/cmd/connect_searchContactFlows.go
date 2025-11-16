package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchContactFlowsCmd = &cobra.Command{
	Use:   "search-contact-flows",
	Short: "Searches the flows in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchContactFlowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_searchContactFlowsCmd).Standalone()

		connect_searchContactFlowsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_searchContactFlowsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_searchContactFlowsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_searchContactFlowsCmd.Flags().String("search-criteria", "", "The search criteria to be used to return flows.")
		connect_searchContactFlowsCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
		connect_searchContactFlowsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_searchContactFlowsCmd)
}
