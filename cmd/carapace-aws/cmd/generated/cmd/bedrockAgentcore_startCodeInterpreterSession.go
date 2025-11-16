package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_startCodeInterpreterSessionCmd = &cobra.Command{
	Use:   "start-code-interpreter-session",
	Short: "Creates and initializes a code interpreter session in Amazon Bedrock.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_startCodeInterpreterSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_startCodeInterpreterSessionCmd).Standalone()

		bedrockAgentcore_startCodeInterpreterSessionCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrockAgentcore_startCodeInterpreterSessionCmd.Flags().String("code-interpreter-identifier", "", "The unique identifier of the code interpreter to use for this session.")
		bedrockAgentcore_startCodeInterpreterSessionCmd.Flags().String("name", "", "The name of the code interpreter session.")
		bedrockAgentcore_startCodeInterpreterSessionCmd.Flags().String("session-timeout-seconds", "", "The time in seconds after which the session automatically terminates if there is no activity.")
		bedrockAgentcore_startCodeInterpreterSessionCmd.Flags().String("trace-id", "", "The trace identifier for request tracking.")
		bedrockAgentcore_startCodeInterpreterSessionCmd.Flags().String("trace-parent", "", "The parent trace information for distributed tracing.")
		bedrockAgentcore_startCodeInterpreterSessionCmd.MarkFlagRequired("code-interpreter-identifier")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_startCodeInterpreterSessionCmd)
}
