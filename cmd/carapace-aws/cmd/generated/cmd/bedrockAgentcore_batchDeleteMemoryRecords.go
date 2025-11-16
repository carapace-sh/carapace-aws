package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_batchDeleteMemoryRecordsCmd = &cobra.Command{
	Use:   "batch-delete-memory-records",
	Short: "Deletes multiple memory records in a single batch operation from the specified memory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_batchDeleteMemoryRecordsCmd).Standalone()

	bedrockAgentcore_batchDeleteMemoryRecordsCmd.Flags().String("memory-id", "", "The unique ID of the memory resource where records will be deleted.")
	bedrockAgentcore_batchDeleteMemoryRecordsCmd.Flags().String("records", "", "A list of memory record deletion inputs to be processed in the batch operation.")
	bedrockAgentcore_batchDeleteMemoryRecordsCmd.MarkFlagRequired("memory-id")
	bedrockAgentcore_batchDeleteMemoryRecordsCmd.MarkFlagRequired("records")
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_batchDeleteMemoryRecordsCmd)
}
