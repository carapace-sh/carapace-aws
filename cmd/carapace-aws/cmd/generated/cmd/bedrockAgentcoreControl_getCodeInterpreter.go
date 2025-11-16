package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_getCodeInterpreterCmd = &cobra.Command{
	Use:   "get-code-interpreter",
	Short: "Gets information about a custom code interpreter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_getCodeInterpreterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_getCodeInterpreterCmd).Standalone()

		bedrockAgentcoreControl_getCodeInterpreterCmd.Flags().String("code-interpreter-id", "", "The unique identifier of the code interpreter to retrieve.")
		bedrockAgentcoreControl_getCodeInterpreterCmd.MarkFlagRequired("code-interpreter-id")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_getCodeInterpreterCmd)
}
