package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_getCodeInterpreterSessionCmd = &cobra.Command{
	Use:   "get-code-interpreter-session",
	Short: "Retrieves detailed information about a specific code interpreter session in Amazon Bedrock.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_getCodeInterpreterSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_getCodeInterpreterSessionCmd).Standalone()

		bedrockAgentcore_getCodeInterpreterSessionCmd.Flags().String("code-interpreter-identifier", "", "The unique identifier of the code interpreter associated with the session.")
		bedrockAgentcore_getCodeInterpreterSessionCmd.Flags().String("session-id", "", "The unique identifier of the code interpreter session to retrieve.")
		bedrockAgentcore_getCodeInterpreterSessionCmd.MarkFlagRequired("code-interpreter-identifier")
		bedrockAgentcore_getCodeInterpreterSessionCmd.MarkFlagRequired("session-id")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_getCodeInterpreterSessionCmd)
}
