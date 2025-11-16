package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listPoliciesCmd = &cobra.Command{
	Use:   "list-policies",
	Short: "Retrieves the list of all policies in an organization of a specified type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listPoliciesCmd).Standalone()

	organizations_listPoliciesCmd.Flags().String("filter", "", "Specifies the type of policy that you want to include in the response.")
	organizations_listPoliciesCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
	organizations_listPoliciesCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	organizations_listPoliciesCmd.MarkFlagRequired("filter")
	organizationsCmd.AddCommand(organizations_listPoliciesCmd)
}
