package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_deleteModelInvocationLoggingConfigurationCmd = &cobra.Command{
	Use:   "delete-model-invocation-logging-configuration",
	Short: "Delete the invocation logging.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_deleteModelInvocationLoggingConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_deleteModelInvocationLoggingConfigurationCmd).Standalone()

	})
	bedrockCmd.AddCommand(bedrock_deleteModelInvocationLoggingConfigurationCmd)
}
