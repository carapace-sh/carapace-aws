package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeAccountPoliciesCmd = &cobra.Command{
	Use:   "describe-account-policies",
	Short: "Returns a list of all CloudWatch Logs account policies in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeAccountPoliciesCmd).Standalone()

	logs_describeAccountPoliciesCmd.Flags().String("account-identifiers", "", "If you are using an account that is set up as a monitoring account for CloudWatch unified cross-account observability, you can use this to specify the account ID of a source account.")
	logs_describeAccountPoliciesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	logs_describeAccountPoliciesCmd.Flags().String("policy-name", "", "Use this parameter to limit the returned policies to only the policy with the name that you specify.")
	logs_describeAccountPoliciesCmd.Flags().String("policy-type", "", "Use this parameter to limit the returned policies to only the policies that match the policy type that you specify.")
	logs_describeAccountPoliciesCmd.MarkFlagRequired("policy-type")
	logsCmd.AddCommand(logs_describeAccountPoliciesCmd)
}
