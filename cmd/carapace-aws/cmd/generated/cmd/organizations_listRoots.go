package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listRootsCmd = &cobra.Command{
	Use:   "list-roots",
	Short: "Lists the roots that are defined in the current organization.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listRootsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_listRootsCmd).Standalone()

		organizations_listRootsCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
		organizations_listRootsCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	})
	organizationsCmd.AddCommand(organizations_listRootsCmd)
}
