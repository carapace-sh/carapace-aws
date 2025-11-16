package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_createSessionCmd = &cobra.Command{
	Use:   "create-session",
	Short: "Creates a session to temporarily store conversations for generative AI (GenAI) applications built with open-source frameworks such as LangGraph and LlamaIndex.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_createSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_createSessionCmd).Standalone()

		bedrockAgentRuntime_createSessionCmd.Flags().String("encryption-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key to use to encrypt the session data.")
		bedrockAgentRuntime_createSessionCmd.Flags().String("session-metadata", "", "A map of key-value pairs containing attributes to be persisted across the session.")
		bedrockAgentRuntime_createSessionCmd.Flags().String("tags", "", "Specify the key-value pairs for the tags that you want to attach to the session.")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_createSessionCmd)
}
