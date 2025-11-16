package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putDataProtectionPolicyCmd = &cobra.Command{
	Use:   "put-data-protection-policy",
	Short: "Creates a data protection policy for the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putDataProtectionPolicyCmd).Standalone()

	logs_putDataProtectionPolicyCmd.Flags().String("log-group-identifier", "", "Specify either the log group name or log group ARN.")
	logs_putDataProtectionPolicyCmd.Flags().String("policy-document", "", "Specify the data protection policy, in JSON.")
	logs_putDataProtectionPolicyCmd.MarkFlagRequired("log-group-identifier")
	logs_putDataProtectionPolicyCmd.MarkFlagRequired("policy-document")
	logsCmd.AddCommand(logs_putDataProtectionPolicyCmd)
}
