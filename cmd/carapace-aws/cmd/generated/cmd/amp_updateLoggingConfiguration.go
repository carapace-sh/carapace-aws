package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_updateLoggingConfigurationCmd = &cobra.Command{
	Use:   "update-logging-configuration",
	Short: "Updates the log group ARN or the workspace ID of the current rules and alerting logging configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_updateLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_updateLoggingConfigurationCmd).Standalone()

		amp_updateLoggingConfigurationCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
		amp_updateLoggingConfigurationCmd.Flags().String("log-group-arn", "", "The ARN of the CloudWatch log group to which the vended log data will be published.")
		amp_updateLoggingConfigurationCmd.Flags().String("workspace-id", "", "The ID of the workspace to update the logging configuration for.")
		amp_updateLoggingConfigurationCmd.MarkFlagRequired("log-group-arn")
		amp_updateLoggingConfigurationCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_updateLoggingConfigurationCmd)
}
