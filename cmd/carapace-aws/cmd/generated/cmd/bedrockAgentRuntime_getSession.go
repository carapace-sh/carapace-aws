package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_getSessionCmd = &cobra.Command{
	Use:   "get-session",
	Short: "Retrieves details about a specific session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_getSessionCmd).Standalone()

	bedrockAgentRuntime_getSessionCmd.Flags().String("session-identifier", "", "A unique identifier for the session to retrieve.")
	bedrockAgentRuntime_getSessionCmd.MarkFlagRequired("session-identifier")
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_getSessionCmd)
}
