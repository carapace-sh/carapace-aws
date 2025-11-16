package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentRuntimeCmd = &cobra.Command{
	Use:   "bedrock-agent-runtime",
	Short: "Contains APIs related to model invocation and querying of knowledge bases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentRuntimeCmd).Standalone()

	rootCmd.AddCommand(bedrockAgentRuntimeCmd)
}
