package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_createLoggingConfigurationCmd = &cobra.Command{
	Use:   "create-logging-configuration",
	Short: "The `CreateLoggingConfiguration` operation creates rules and alerting logging configuration for the workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_createLoggingConfigurationCmd).Standalone()

	amp_createLoggingConfigurationCmd.Flags().String("client-token", "", "A unique identifier that you can provide to ensure the idempotency of the request.")
	amp_createLoggingConfigurationCmd.Flags().String("log-group-arn", "", "The ARN of the CloudWatch log group to which the vended log data will be published.")
	amp_createLoggingConfigurationCmd.Flags().String("workspace-id", "", "The ID of the workspace to create the logging configuration for.")
	amp_createLoggingConfigurationCmd.MarkFlagRequired("log-group-arn")
	amp_createLoggingConfigurationCmd.MarkFlagRequired("workspace-id")
	ampCmd.AddCommand(amp_createLoggingConfigurationCmd)
}
