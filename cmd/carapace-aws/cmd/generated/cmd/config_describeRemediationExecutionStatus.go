package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_describeRemediationExecutionStatusCmd = &cobra.Command{
	Use:   "describe-remediation-execution-status",
	Short: "Provides a detailed view of a Remediation Execution for a set of resources including state, timestamps for when steps for the remediation execution occur, and any error messages for steps that have failed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_describeRemediationExecutionStatusCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_describeRemediationExecutionStatusCmd).Standalone()

		config_describeRemediationExecutionStatusCmd.Flags().String("config-rule-name", "", "The name of the Config rule.")
		config_describeRemediationExecutionStatusCmd.Flags().String("limit", "", "The maximum number of RemediationExecutionStatuses returned on each page.")
		config_describeRemediationExecutionStatusCmd.Flags().String("next-token", "", "The `nextToken` string returned on a previous page that you use to get the next page of results in a paginated response.")
		config_describeRemediationExecutionStatusCmd.Flags().String("resource-keys", "", "A list of resource keys to be processed with the current request.")
		config_describeRemediationExecutionStatusCmd.MarkFlagRequired("config-rule-name")
	})
	configCmd.AddCommand(config_describeRemediationExecutionStatusCmd)
}
