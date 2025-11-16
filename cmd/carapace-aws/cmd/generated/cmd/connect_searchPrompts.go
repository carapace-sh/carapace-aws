package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchPromptsCmd = &cobra.Command{
	Use:   "search-prompts",
	Short: "Searches prompts in an Amazon Connect instance, with optional filtering.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchPromptsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_searchPromptsCmd).Standalone()

		connect_searchPromptsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_searchPromptsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_searchPromptsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_searchPromptsCmd.Flags().String("search-criteria", "", "The search criteria to be used to return prompts.")
		connect_searchPromptsCmd.Flags().String("search-filter", "", "Filters to be applied to search results.")
		connect_searchPromptsCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_searchPromptsCmd)
}
