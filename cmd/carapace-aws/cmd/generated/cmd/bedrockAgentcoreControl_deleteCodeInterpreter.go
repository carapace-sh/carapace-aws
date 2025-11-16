package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_deleteCodeInterpreterCmd = &cobra.Command{
	Use:   "delete-code-interpreter",
	Short: "Deletes a custom code interpreter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_deleteCodeInterpreterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreControl_deleteCodeInterpreterCmd).Standalone()

		bedrockAgentcoreControl_deleteCodeInterpreterCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
		bedrockAgentcoreControl_deleteCodeInterpreterCmd.Flags().String("code-interpreter-id", "", "The unique identifier of the code interpreter to delete.")
		bedrockAgentcoreControl_deleteCodeInterpreterCmd.MarkFlagRequired("code-interpreter-id")
	})
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_deleteCodeInterpreterCmd)
}
