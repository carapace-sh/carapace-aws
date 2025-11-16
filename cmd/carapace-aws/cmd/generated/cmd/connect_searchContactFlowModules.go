package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchContactFlowModulesCmd = &cobra.Command{
	Use:   "search-contact-flow-modules",
	Short: "Searches the flow modules in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchContactFlowModulesCmd).Standalone()

	connect_searchContactFlowModulesCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_searchContactFlowModulesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_searchContactFlowModulesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_searchContactFlowModulesCmd.Flags().String("search-criteria", "", "The search criteria to be used to return flow modules.")
	connect_searchContactFlowModulesCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
	connect_searchContactFlowModulesCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_searchContactFlowModulesCmd)
}
