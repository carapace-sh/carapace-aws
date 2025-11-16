package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listChildrenCmd = &cobra.Command{
	Use:   "list-children",
	Short: "Lists all of the organizational units (OUs) or accounts that are contained in the specified parent OU or root.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listChildrenCmd).Standalone()

	organizations_listChildrenCmd.Flags().String("child-type", "", "Filters the output to include only the specified child type.")
	organizations_listChildrenCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
	organizations_listChildrenCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	organizations_listChildrenCmd.Flags().String("parent-id", "", "The unique identifier (ID) for the parent root or OU whose children you want to list.")
	organizations_listChildrenCmd.MarkFlagRequired("child-type")
	organizations_listChildrenCmd.MarkFlagRequired("parent-id")
	organizationsCmd.AddCommand(organizations_listChildrenCmd)
}
