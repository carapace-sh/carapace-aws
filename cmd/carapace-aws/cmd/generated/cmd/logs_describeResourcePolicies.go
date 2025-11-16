package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeResourcePoliciesCmd = &cobra.Command{
	Use:   "describe-resource-policies",
	Short: "Lists the resource policies in this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeResourcePoliciesCmd).Standalone()

	logs_describeResourcePoliciesCmd.Flags().String("limit", "", "The maximum number of resource policies to be displayed with one call of this API.")
	logs_describeResourcePoliciesCmd.Flags().String("next-token", "", "")
	logs_describeResourcePoliciesCmd.Flags().String("policy-scope", "", "Specifies the scope of the resource policy.")
	logs_describeResourcePoliciesCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch Logs resource for which to query the resource policy.")
	logsCmd.AddCommand(logs_describeResourcePoliciesCmd)
}
