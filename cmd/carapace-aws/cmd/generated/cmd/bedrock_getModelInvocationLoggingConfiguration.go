package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_getModelInvocationLoggingConfigurationCmd = &cobra.Command{
	Use:   "get-model-invocation-logging-configuration",
	Short: "Get the current configuration values for model invocation logging.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_getModelInvocationLoggingConfigurationCmd).Standalone()

	bedrockCmd.AddCommand(bedrock_getModelInvocationLoggingConfigurationCmd)
}
