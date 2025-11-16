package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listMembersCmd = &cobra.Command{
	Use:   "list-members",
	Short: "Lists all members within a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listMembersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_listMembersCmd).Standalone()

		cleanrooms_listMembersCmd.Flags().String("collaboration-identifier", "", "The identifier of the collaboration in which the members are listed.")
		cleanrooms_listMembersCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
		cleanrooms_listMembersCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		cleanrooms_listMembersCmd.MarkFlagRequired("collaboration-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_listMembersCmd)
}
