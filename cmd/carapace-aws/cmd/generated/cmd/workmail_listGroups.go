package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "Returns summaries of the organization's groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_listGroupsCmd).Standalone()

		workmail_listGroupsCmd.Flags().String("filters", "", "Limit the search results based on the filter criteria.")
		workmail_listGroupsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		workmail_listGroupsCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		workmail_listGroupsCmd.Flags().String("organization-id", "", "The identifier for the organization under which the groups exist.")
		workmail_listGroupsCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_listGroupsCmd)
}
