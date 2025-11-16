package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_startRemediationExecutionCmd = &cobra.Command{
	Use:   "start-remediation-execution",
	Short: "Runs an on-demand remediation for the specified Config rules against the last known remediation configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_startRemediationExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(config_startRemediationExecutionCmd).Standalone()

		config_startRemediationExecutionCmd.Flags().String("config-rule-name", "", "The list of names of Config rules that you want to run remediation execution for.")
		config_startRemediationExecutionCmd.Flags().String("resource-keys", "", "A list of resource keys to be processed with the current request.")
		config_startRemediationExecutionCmd.MarkFlagRequired("config-rule-name")
		config_startRemediationExecutionCmd.MarkFlagRequired("resource-keys")
	})
	configCmd.AddCommand(config_startRemediationExecutionCmd)
}
