package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_deleteSessionCmd = &cobra.Command{
	Use:   "delete-session",
	Short: "Deletes a session that you ended.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_deleteSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_deleteSessionCmd).Standalone()

		bedrockAgentRuntime_deleteSessionCmd.Flags().String("session-identifier", "", "The unique identifier for the session to be deleted.")
		bedrockAgentRuntime_deleteSessionCmd.MarkFlagRequired("session-identifier")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_deleteSessionCmd)
}
