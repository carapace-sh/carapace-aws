package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgentcore_batchCreateMemoryRecordsCmd = &cobra.Command{
	Use:   "batch-create-memory-records",
	Short: "Creates multiple memory records in a single batch operation for the specified memory with custom content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgentcore_batchCreateMemoryRecordsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgentcore_batchCreateMemoryRecordsCmd).Standalone()

		bedrockAgentcore_batchCreateMemoryRecordsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotent processing of the batch request.")
		bedrockAgentcore_batchCreateMemoryRecordsCmd.Flags().String("memory-id", "", "The unique ID of the memory resource where records will be created.")
		bedrockAgentcore_batchCreateMemoryRecordsCmd.Flags().String("records", "", "A list of memory record creation inputs to be processed in the batch operation.")
		bedrockAgentcore_batchCreateMemoryRecordsCmd.MarkFlagRequired("memory-id")
		bedrockAgentcore_batchCreateMemoryRecordsCmd.MarkFlagRequired("records")
	})
	bedrockAgentcoreCmd.AddCommand(bedrockAgentcore_batchCreateMemoryRecordsCmd)
}
