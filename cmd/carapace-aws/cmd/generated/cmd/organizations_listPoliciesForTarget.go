package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listPoliciesForTargetCmd = &cobra.Command{
	Use:   "list-policies-for-target",
	Short: "Lists the policies that are directly attached to the specified target root, organizational unit (OU), or account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listPoliciesForTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(organizations_listPoliciesForTargetCmd).Standalone()

		organizations_listPoliciesForTargetCmd.Flags().String("filter", "", "The type of policy that you want to include in the returned list.")
		organizations_listPoliciesForTargetCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
		organizations_listPoliciesForTargetCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
		organizations_listPoliciesForTargetCmd.Flags().String("target-id", "", "The unique identifier (ID) of the root, organizational unit, or account whose policies you want to list.")
		organizations_listPoliciesForTargetCmd.MarkFlagRequired("filter")
		organizations_listPoliciesForTargetCmd.MarkFlagRequired("target-id")
	})
	organizationsCmd.AddCommand(organizations_listPoliciesForTargetCmd)
}
