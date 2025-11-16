package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteAccountPolicyCmd = &cobra.Command{
	Use:   "delete-account-policy",
	Short: "Deletes a CloudWatch Logs account policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteAccountPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_deleteAccountPolicyCmd).Standalone()

		logs_deleteAccountPolicyCmd.Flags().String("policy-name", "", "The name of the policy to delete.")
		logs_deleteAccountPolicyCmd.Flags().String("policy-type", "", "The type of policy to delete.")
		logs_deleteAccountPolicyCmd.MarkFlagRequired("policy-name")
		logs_deleteAccountPolicyCmd.MarkFlagRequired("policy-type")
	})
	logsCmd.AddCommand(logs_deleteAccountPolicyCmd)
}
