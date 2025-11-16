package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_retrieveMemoryRecordsCmd = &cobra.Command{
	Use:   "retrieve-memory-records",
	Short: "Searches for and retrieves memory records from an AgentCore Memory resource based on specified search criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_retrieveMemoryRecordsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_retrieveMemoryRecordsCmd).Standalone()

		bedrockAgentcore_retrieveMemoryRecordsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		bedrockAgentcore_retrieveMemoryRecordsCmd.Flags().String("memory-id", "", "The identifier of the AgentCore Memory resource from which to retrieve memory records.")
		bedrockAgentcore_retrieveMemoryRecordsCmd.Flags().String("namespace", "", "The namespace to filter memory records by.")
		bedrockAgentcore_retrieveMemoryRecordsCmd.Flags().String("next-token", "", "The token for the next set of results.")
		bedrockAgentcore_retrieveMemoryRecordsCmd.Flags().String("search-criteria", "", "The search criteria to use for finding relevant memory records.")
		bedrockAgentcore_retrieveMemoryRecordsCmd.MarkFlagRequired("memory-id")
		bedrockAgentcore_retrieveMemoryRecordsCmd.MarkFlagRequired("namespace")
		bedrockAgentcore_retrieveMemoryRecordsCmd.MarkFlagRequired("search-criteria")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_retrieveMemoryRecordsCmd)
}
