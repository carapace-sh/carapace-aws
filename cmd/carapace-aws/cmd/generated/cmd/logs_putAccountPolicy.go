package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putAccountPolicyCmd = &cobra.Command{
	Use:   "put-account-policy",
	Short: "Creates an account-level data protection policy, subscription filter policy, field index policy, transformer policy, or metric extraction policy that applies to all log groups or a subset of log groups in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putAccountPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_putAccountPolicyCmd).Standalone()

		logs_putAccountPolicyCmd.Flags().String("policy-document", "", "Specify the policy, in JSON.")
		logs_putAccountPolicyCmd.Flags().String("policy-name", "", "A name for the policy.")
		logs_putAccountPolicyCmd.Flags().String("policy-type", "", "The type of policy that you're creating or updating.")
		logs_putAccountPolicyCmd.Flags().String("scope", "", "Currently the only valid value for this parameter is `ALL`, which specifies that the data protection policy applies to all log groups in the account.")
		logs_putAccountPolicyCmd.Flags().String("selection-criteria", "", "Use this parameter to apply the new policy to a subset of log groups in the account.")
		logs_putAccountPolicyCmd.MarkFlagRequired("policy-document")
		logs_putAccountPolicyCmd.MarkFlagRequired("policy-name")
		logs_putAccountPolicyCmd.MarkFlagRequired("policy-type")
	})
	logsCmd.AddCommand(logs_putAccountPolicyCmd)
}
