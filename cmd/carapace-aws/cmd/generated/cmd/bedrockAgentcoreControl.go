package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControlCmd = &cobra.Command{
	Use:   "bedrock-agentcore-control",
	Short: "Welcome to the Amazon Bedrock AgentCore Control plane API reference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControlCmd).Standalone()

	rootCmd.AddCommand(bedrockAgentcoreControlCmd)
}
