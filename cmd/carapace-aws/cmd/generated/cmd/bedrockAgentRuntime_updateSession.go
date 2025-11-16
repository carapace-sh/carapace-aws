package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_updateSessionCmd = &cobra.Command{
	Use:   "update-session",
	Short: "Updates the metadata or encryption settings of a session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_updateSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_updateSessionCmd).Standalone()

		bedrockAgentRuntime_updateSessionCmd.Flags().String("session-identifier", "", "The unique identifier of the session to modify.")
		bedrockAgentRuntime_updateSessionCmd.Flags().String("session-metadata", "", "A map of key-value pairs containing attributes to be persisted across the session.")
		bedrockAgentRuntime_updateSessionCmd.MarkFlagRequired("session-identifier")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_updateSessionCmd)
}
