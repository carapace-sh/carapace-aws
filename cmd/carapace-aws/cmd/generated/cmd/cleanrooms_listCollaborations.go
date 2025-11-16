package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listCollaborationsCmd = &cobra.Command{
	Use:   "list-collaborations",
	Short: "Lists collaborations the caller owns, is active in, or has been invited to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listCollaborationsCmd).Standalone()

	cleanrooms_listCollaborationsCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
	cleanrooms_listCollaborationsCmd.Flags().String("member-status", "", "The caller's status in a collaboration.")
	cleanrooms_listCollaborationsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	cleanroomsCmd.AddCommand(cleanrooms_listCollaborationsCmd)
}
