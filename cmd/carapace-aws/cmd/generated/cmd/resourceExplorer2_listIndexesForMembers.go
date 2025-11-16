package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_listIndexesForMembersCmd = &cobra.Command{
	Use:   "list-indexes-for-members",
	Short: "Retrieves a list of a member's indexes in all Amazon Web Services Regions that are currently collecting resource information for Amazon Web Services Resource Explorer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_listIndexesForMembersCmd).Standalone()

	resourceExplorer2_listIndexesForMembersCmd.Flags().String("account-id-list", "", "The account IDs will limit the output to only indexes from these accounts.")
	resourceExplorer2_listIndexesForMembersCmd.Flags().String("max-results", "", "The maximum number of results that you want included on each page of the response.")
	resourceExplorer2_listIndexesForMembersCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	resourceExplorer2_listIndexesForMembersCmd.MarkFlagRequired("account-id-list")
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_listIndexesForMembersCmd)
}
