package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_listMemoryRecordsCmd = &cobra.Command{
	Use:   "list-memory-records",
	Short: "Lists memory records in an AgentCore Memory resource based on specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_listMemoryRecordsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_listMemoryRecordsCmd).Standalone()

		bedrockAgentcore_listMemoryRecordsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		bedrockAgentcore_listMemoryRecordsCmd.Flags().String("memory-id", "", "The identifier of the AgentCore Memory resource for which to list memory records.")
		bedrockAgentcore_listMemoryRecordsCmd.Flags().String("memory-strategy-id", "", "The memory strategy identifier to filter memory records by.")
		bedrockAgentcore_listMemoryRecordsCmd.Flags().String("namespace", "", "The namespace to filter memory records by.")
		bedrockAgentcore_listMemoryRecordsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		bedrockAgentcore_listMemoryRecordsCmd.MarkFlagRequired("memory-id")
		bedrockAgentcore_listMemoryRecordsCmd.MarkFlagRequired("namespace")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_listMemoryRecordsCmd)
}
