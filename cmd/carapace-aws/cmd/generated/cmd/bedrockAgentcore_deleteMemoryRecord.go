package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_deleteMemoryRecordCmd = &cobra.Command{
	Use:   "delete-memory-record",
	Short: "Deletes a memory record from an AgentCore Memory resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_deleteMemoryRecordCmd).Standalone()

	bedrockAgentcore_deleteMemoryRecordCmd.Flags().String("memory-id", "", "The identifier of the AgentCore Memory resource from which to delete the memory record.")
	bedrockAgentcore_deleteMemoryRecordCmd.Flags().String("memory-record-id", "", "The identifier of the memory record to delete.")
	bedrockAgentcore_deleteMemoryRecordCmd.MarkFlagRequired("memory-id")
	bedrockAgentcore_deleteMemoryRecordCmd.MarkFlagRequired("memory-record-id")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_deleteMemoryRecordCmd)
}
