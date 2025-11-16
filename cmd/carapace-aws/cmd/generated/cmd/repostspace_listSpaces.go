package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repostspace_listSpacesCmd = &cobra.Command{
	Use:   "list-spaces",
	Short: "Returns a list of AWS re:Post Private private re:Posts in the account with some information about each private re:Post.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repostspace_listSpacesCmd).Standalone()

	repostspace_listSpacesCmd.Flags().String("max-results", "", "The maximum number of private re:Posts to include in the results.")
	repostspace_listSpacesCmd.Flags().String("next-token", "", "The token for the next set of private re:Posts to return.")
	repostspaceCmd.AddCommand(repostspace_listSpacesCmd)
}
