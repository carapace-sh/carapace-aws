package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_getMemoryRecordCmd = &cobra.Command{
	Use:   "get-memory-record",
	Short: "Retrieves a specific memory record from an AgentCore Memory resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_getMemoryRecordCmd).Standalone()

	bedrockAgentcore_getMemoryRecordCmd.Flags().String("memory-id", "", "The identifier of the AgentCore Memory resource containing the memory record.")
	bedrockAgentcore_getMemoryRecordCmd.Flags().String("memory-record-id", "", "The identifier of the memory record to retrieve.")
	bedrockAgentcore_getMemoryRecordCmd.MarkFlagRequired("memory-id")
	bedrockAgentcore_getMemoryRecordCmd.MarkFlagRequired("memory-record-id")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_getMemoryRecordCmd)
}
