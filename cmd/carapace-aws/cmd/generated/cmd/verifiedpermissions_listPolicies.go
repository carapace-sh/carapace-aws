package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifiedpermissions_listPoliciesCmd = &cobra.Command{
	Use:   "list-policies",
	Short: "Returns a paginated list of all policies stored in the specified policy store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifiedpermissions_listPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(verifiedpermissions_listPoliciesCmd).Standalone()

		verifiedpermissions_listPoliciesCmd.Flags().String("filter", "", "Specifies a filter that limits the response to only policies that match the specified criteria.")
		verifiedpermissions_listPoliciesCmd.Flags().String("max-results", "", "Specifies the total number of results that you want included in each response.")
		verifiedpermissions_listPoliciesCmd.Flags().String("next-token", "", "Specifies that you want to receive the next page of results.")
		verifiedpermissions_listPoliciesCmd.Flags().String("policy-store-id", "", "Specifies the ID of the policy store you want to list policies from.")
		verifiedpermissions_listPoliciesCmd.MarkFlagRequired("policy-store-id")
	})
	verifiedpermissionsCmd.AddCommand(verifiedpermissions_listPoliciesCmd)
}
