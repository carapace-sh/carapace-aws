package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_getAgentCardCmd = &cobra.Command{
	Use:   "get-agent-card",
	Short: "Retrieves the A2A agent card associated with an AgentCore Runtime agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_getAgentCardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_getAgentCardCmd).Standalone()

		bedrockAgentcore_getAgentCardCmd.Flags().String("agent-runtime-arn", "", "The ARN of the AgentCore Runtime agent for which you want to get the A2A agent card.")
		bedrockAgentcore_getAgentCardCmd.Flags().String("qualifier", "", "Optional qualifier to specify an agent alias, such as `prod`code&gt; or `dev`.")
		bedrockAgentcore_getAgentCardCmd.Flags().String("runtime-session-id", "", "The session ID that the AgentCore Runtime agent is using.")
		bedrockAgentcore_getAgentCardCmd.MarkFlagRequired("agent-runtime-arn")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_getAgentCardCmd)
}
