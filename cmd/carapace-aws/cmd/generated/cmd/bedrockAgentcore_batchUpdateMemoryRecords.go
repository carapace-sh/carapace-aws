package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_batchUpdateMemoryRecordsCmd = &cobra.Command{
	Use:   "batch-update-memory-records",
	Short: "Updates multiple memory records with custom content in a single batch operation within the specified memory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_batchUpdateMemoryRecordsCmd).Standalone()

	bedrockAgentcore_batchUpdateMemoryRecordsCmd.Flags().String("memory-id", "", "The unique ID of the memory resource where records will be updated.")
	bedrockAgentcore_batchUpdateMemoryRecordsCmd.Flags().String("records", "", "A list of memory record update inputs to be processed in the batch operation.")
	bedrockAgentcore_batchUpdateMemoryRecordsCmd.MarkFlagRequired("memory-id")
	bedrockAgentcore_batchUpdateMemoryRecordsCmd.MarkFlagRequired("records")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_batchUpdateMemoryRecordsCmd)
}
