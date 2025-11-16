package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_stopCodeInterpreterSessionCmd = &cobra.Command{
	Use:   "stop-code-interpreter-session",
	Short: "Terminates an active code interpreter session in Amazon Bedrock.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_stopCodeInterpreterSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_stopCodeInterpreterSessionCmd).Standalone()

		bedrockAgentcore_stopCodeInterpreterSessionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrockAgentcore_stopCodeInterpreterSessionCmd.Flags().String("code-interpreter-identifier", "", "The unique identifier of the code interpreter associated with the session.")
		bedrockAgentcore_stopCodeInterpreterSessionCmd.Flags().String("session-id", "", "The unique identifier of the code interpreter session to stop.")
		bedrockAgentcore_stopCodeInterpreterSessionCmd.Flags().String("trace-id", "", "The trace identifier for request tracking.")
		bedrockAgentcore_stopCodeInterpreterSessionCmd.Flags().String("trace-parent", "", "The parent trace information for distributed tracing.")
		bedrockAgentcore_stopCodeInterpreterSessionCmd.MarkFlagRequired("code-interpreter-identifier")
		bedrockAgentcore_stopCodeInterpreterSessionCmd.MarkFlagRequired("session-id")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_stopCodeInterpreterSessionCmd)
}
