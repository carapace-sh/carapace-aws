package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listParentsCmd = &cobra.Command{
	Use:   "list-parents",
	Short: "Lists the root or organizational units (OUs) that serve as the immediate parent of the specified child OU or account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listParentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_listParentsCmd).Standalone()

		organizations_listParentsCmd.Flags().String("child-id", "", "The unique identifier (ID) of the OU or account whose parent containers you want to list.")
		organizations_listParentsCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
		organizations_listParentsCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
		organizations_listParentsCmd.MarkFlagRequired("child-id")
	})
	organizationsCmd.AddCommand(organizations_listParentsCmd)
}
