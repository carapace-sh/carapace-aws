package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listAllowListsCmd = &cobra.Command{
	Use:   "list-allow-lists",
	Short: "Retrieves a subset of information about all the allow lists for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listAllowListsCmd).Standalone()

	macie2_listAllowListsCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of a paginated response.")
	macie2_listAllowListsCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	macie2Cmd.AddCommand(macie2_listAllowListsCmd)
}
