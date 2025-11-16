package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentCmd = &cobra.Command{
	Use:   "bedrock-agent",
	Short: "Describes the API operations for creating and managing Amazon Bedrock agents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentCmd).Standalone()

	rootCmd.AddCommand(bedrockAgentCmd)
}
