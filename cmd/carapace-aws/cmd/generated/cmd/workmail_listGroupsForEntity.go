package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmail_listGroupsForEntityCmd = &cobra.Command{
	Use:   "list-groups-for-entity",
	Short: "Returns all the groups to which an entity belongs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmail_listGroupsForEntityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmail_listGroupsForEntityCmd).Standalone()

		workmail_listGroupsForEntityCmd.Flags().String("entity-id", "", "The identifier for the entity.")
		workmail_listGroupsForEntityCmd.Flags().String("filters", "", "Limit the search results based on the filter criteria.")
		workmail_listGroupsForEntityCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		workmail_listGroupsForEntityCmd.Flags().String("next-token", "", "The token to use to retrieve the next page of results.")
		workmail_listGroupsForEntityCmd.Flags().String("organization-id", "", "The identifier for the organization under which the entity exists.")
		workmail_listGroupsForEntityCmd.MarkFlagRequired("entity-id")
		workmail_listGroupsForEntityCmd.MarkFlagRequired("organization-id")
	})
	workmailCmd.AddCommand(workmail_listGroupsForEntityCmd)
}
