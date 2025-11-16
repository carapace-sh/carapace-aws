package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreCmd = &cobra.Command{
	Use:   "bedrock-agentcore",
	Short: "Amazon Bedrock AgentCore is in preview release and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcoreCmd).Standalone()

	})
	rootCmd.AddCommand(bedrockAgentcoreCmd)
}
