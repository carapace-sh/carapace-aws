package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntime_endSessionCmd = &cobra.Command{
	Use:   "end-session",
	Short: "Ends the session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntime_endSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentRuntime_endSessionCmd).Standalone()

		bedrockAgentRuntime_endSessionCmd.Flags().String("session-identifier", "", "The unique identifier for the session to end.")
		bedrockAgentRuntime_endSessionCmd.MarkFlagRequired("session-identifier")
	})
	bedrockAgentRuntimeCmd.AddCommand(bedrockAgentRuntime_endSessionCmd)
}
