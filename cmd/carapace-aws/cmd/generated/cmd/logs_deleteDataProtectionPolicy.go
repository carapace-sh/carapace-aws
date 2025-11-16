package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteDataProtectionPolicyCmd = &cobra.Command{
	Use:   "delete-data-protection-policy",
	Short: "Deletes the data protection policy from the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteDataProtectionPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_deleteDataProtectionPolicyCmd).Standalone()

		logs_deleteDataProtectionPolicyCmd.Flags().String("log-group-identifier", "", "The name or ARN of the log group that you want to delete the data protection policy for.")
		logs_deleteDataProtectionPolicyCmd.MarkFlagRequired("log-group-identifier")
	})
	logsCmd.AddCommand(logs_deleteDataProtectionPolicyCmd)
}
