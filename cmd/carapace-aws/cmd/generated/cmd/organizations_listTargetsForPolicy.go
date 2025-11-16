package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizations_listTargetsForPolicyCmd = &cobra.Command{
	Use:   "list-targets-for-policy",
	Short: "Lists all the roots, organizational units (OUs), and accounts that the specified policy is attached to.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizations_listTargetsForPolicyCmd).Standalone()

	organizations_listTargetsForPolicyCmd.Flags().String("max-results", "", "The total number of results that you want included on each page of the response.")
	organizations_listTargetsForPolicyCmd.Flags().String("next-token", "", "The parameter for receiving additional results if you receive a `NextToken` response in a previous request.")
	organizations_listTargetsForPolicyCmd.Flags().String("policy-id", "", "The unique identifier (ID) of the policy whose attachments you want to know.")
	organizations_listTargetsForPolicyCmd.MarkFlagRequired("policy-id")
	organizationsCmd.AddCommand(organizations_listTargetsForPolicyCmd)
}
