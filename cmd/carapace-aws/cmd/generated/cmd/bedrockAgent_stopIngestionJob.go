package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_stopIngestionJobCmd = &cobra.Command{
	Use:   "stop-ingestion-job",
	Short: "Stops a currently running data ingestion job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_stopIngestionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_stopIngestionJobCmd).Standalone()

		bedrockAgent_stopIngestionJobCmd.Flags().String("data-source-id", "", "The unique identifier of the data source for the data ingestion job you want to stop.")
		bedrockAgent_stopIngestionJobCmd.Flags().String("ingestion-job-id", "", "The unique identifier of the data ingestion job you want to stop.")
		bedrockAgent_stopIngestionJobCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base for the data ingestion job you want to stop.")
		bedrockAgent_stopIngestionJobCmd.MarkFlagRequired("data-source-id")
		bedrockAgent_stopIngestionJobCmd.MarkFlagRequired("ingestion-job-id")
		bedrockAgent_stopIngestionJobCmd.MarkFlagRequired("knowledge-base-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_stopIngestionJobCmd)
}
