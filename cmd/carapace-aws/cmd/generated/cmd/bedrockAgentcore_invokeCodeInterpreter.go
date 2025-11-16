package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_invokeCodeInterpreterCmd = &cobra.Command{
	Use:   "invoke-code-interpreter",
	Short: "Executes code within an active code interpreter session in Amazon Bedrock.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_invokeCodeInterpreterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_invokeCodeInterpreterCmd).Standalone()

		bedrockAgentcore_invokeCodeInterpreterCmd.Flags().String("arguments", "", "The arguments for the code interpreter.")
		bedrockAgentcore_invokeCodeInterpreterCmd.Flags().String("code-interpreter-identifier", "", "The unique identifier of the code interpreter associated with the session.")
		bedrockAgentcore_invokeCodeInterpreterCmd.Flags().String("name", "", "The name of the code interpreter to invoke.")
		bedrockAgentcore_invokeCodeInterpreterCmd.Flags().String("session-id", "", "The unique identifier of the code interpreter session to use.")
		bedrockAgentcore_invokeCodeInterpreterCmd.Flags().String("trace-id", "", "The trace identifier for request tracking.")
		bedrockAgentcore_invokeCodeInterpreterCmd.Flags().String("trace-parent", "", "The parent trace information for distributed tracing.")
		bedrockAgentcore_invokeCodeInterpreterCmd.MarkFlagRequired("code-interpreter-identifier")
		bedrockAgentcore_invokeCodeInterpreterCmd.MarkFlagRequired("name")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_invokeCodeInterpreterCmd)
}
