package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_putModelInvocationLoggingConfigurationCmd = &cobra.Command{
	Use:   "put-model-invocation-logging-configuration",
	Short: "Set the configuration values for model invocation logging.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_putModelInvocationLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_putModelInvocationLoggingConfigurationCmd).Standalone()

		bedrock_putModelInvocationLoggingConfigurationCmd.Flags().String("logging-config", "", "The logging configuration values to set.")
		bedrock_putModelInvocationLoggingConfigurationCmd.MarkFlagRequired("logging-config")
	})
	bedrockCmd.AddCommand(bedrock_putModelInvocationLoggingConfigurationCmd)
}
