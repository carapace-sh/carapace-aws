package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcoreControl_listMemoriesCmd = &cobra.Command{
	Use:   "list-memories",
	Short: "Lists the available Amazon Bedrock AgentCore Memory resources in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcoreControl_listMemoriesCmd).Standalone()

	bedrockAgentcoreControl_listMemoriesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
	bedrockAgentcoreControl_listMemoriesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	bedrockAgentcoreControlCmd.AddCommand(bedrockAgentcoreControl_listMemoriesCmd)
}
