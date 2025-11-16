package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createPromptRouterCmd = &cobra.Command{
	Use:   "create-prompt-router",
	Short: "Creates a prompt router that manages the routing of requests between multiple foundation models based on the routing criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createPromptRouterCmd).Standalone()

	bedrock_createPromptRouterCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure idempotency of your requests.")
	bedrock_createPromptRouterCmd.Flags().String("description", "", "An optional description of the prompt router to help identify its purpose.")
	bedrock_createPromptRouterCmd.Flags().String("fallback-model", "", "The default model to use when the routing criteria is not met.")
	bedrock_createPromptRouterCmd.Flags().String("models", "", "A list of foundation models that the prompt router can route requests to.")
	bedrock_createPromptRouterCmd.Flags().String("prompt-router-name", "", "The name of the prompt router.")
	bedrock_createPromptRouterCmd.Flags().String("routing-criteria", "", "The criteria, which is the response quality difference, used to determine how incoming requests are routed to different models.")
	bedrock_createPromptRouterCmd.Flags().String("tags", "", "An array of key-value pairs to apply to this resource as tags.")
	bedrock_createPromptRouterCmd.MarkFlagRequired("fallback-model")
	bedrock_createPromptRouterCmd.MarkFlagRequired("models")
	bedrock_createPromptRouterCmd.MarkFlagRequired("prompt-router-name")
	bedrock_createPromptRouterCmd.MarkFlagRequired("routing-criteria")
	bedrockCmd.AddCommand(bedrock_createPromptRouterCmd)
}
