package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listMembershipsCmd = &cobra.Command{
	Use:   "list-memberships",
	Short: "Lists all memberships resources within the caller's account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listMembershipsCmd).Standalone()

	cleanrooms_listMembershipsCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
	cleanrooms_listMembershipsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	cleanrooms_listMembershipsCmd.Flags().String("status", "", "A filter which will return only memberships in the specified status.")
	cleanroomsCmd.AddCommand(cleanrooms_listMembershipsCmd)
}
