package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_stopRuntimeSessionCmd = &cobra.Command{
	Use:   "stop-runtime-session",
	Short: "Stops a session that is running in an running AgentCore Runtime agent.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_stopRuntimeSessionCmd).Standalone()

	bedrockAgentcore_stopRuntimeSessionCmd.Flags().String("agent-runtime-arn", "", "The ARN of the agent that contains the session that you want to stop.")
	bedrockAgentcore_stopRuntimeSessionCmd.Flags().String("client-token", "", "Idempotent token used to identify the request.")
	bedrockAgentcore_stopRuntimeSessionCmd.Flags().String("qualifier", "", "Optional qualifier to specify an agent alias, such as `prod`code&gt; or `dev`.")
	bedrockAgentcore_stopRuntimeSessionCmd.Flags().String("runtime-session-id", "", "The ID of the session that you want to stop.")
	bedrockAgentcore_stopRuntimeSessionCmd.MarkFlagRequired("agent-runtime-arn")
	bedrockAgentcore_stopRuntimeSessionCmd.MarkFlagRequired("runtime-session-id")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_stopRuntimeSessionCmd)
}
