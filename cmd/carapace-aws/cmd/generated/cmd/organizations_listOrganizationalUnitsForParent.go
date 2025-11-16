package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listOrganizationalUnitsForParentCmd = &cobra.Command{
	Use:   "list-organizational-units-for-parent",
	Short: "Lists the organizational units (OUs) in a parent organizational unit or root.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listOrganizationalUnitsForParentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_listOrganizationalUnitsForParentCmd).Standalone()

		organizations_listOrganizationalUnitsForParentCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
		organizations_listOrganizationalUnitsForParentCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
		organizations_listOrganizationalUnitsForParentCmd.Flags().String("parent-id", "", "The unique identifier (ID) of the root or OU whose child OUs you want to list.")
		organizations_listOrganizationalUnitsForParentCmd.MarkFlagRequired("parent-id")
	})
	organizationsCmd.AddCommand(organizations_listOrganizationalUnitsForParentCmd)
}
