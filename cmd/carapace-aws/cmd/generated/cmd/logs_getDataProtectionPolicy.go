package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getDataProtectionPolicyCmd = &cobra.Command{
	Use:   "get-data-protection-policy",
	Short: "Returns information about a log group data protection policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getDataProtectionPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_getDataProtectionPolicyCmd).Standalone()

		logs_getDataProtectionPolicyCmd.Flags().String("log-group-identifier", "", "The name or ARN of the log group that contains the data protection policy that you want to see.")
		logs_getDataProtectionPolicyCmd.MarkFlagRequired("log-group-identifier")
	})
	logsCmd.AddCommand(logs_getDataProtectionPolicyCmd)
}
