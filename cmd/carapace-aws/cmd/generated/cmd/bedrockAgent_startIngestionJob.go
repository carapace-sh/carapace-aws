package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_startIngestionJobCmd = &cobra.Command{
	Use:   "start-ingestion-job",
	Short: "Begins a data ingestion job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_startIngestionJobCmd).Standalone()

	bedrockAgent_startIngestionJobCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
	bedrockAgent_startIngestionJobCmd.Flags().String("data-source-id", "", "The unique identifier of the data source you want to ingest into your knowledge base.")
	bedrockAgent_startIngestionJobCmd.Flags().String("description", "", "A description of the data ingestion job.")
	bedrockAgent_startIngestionJobCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base for the data ingestion job.")
	bedrockAgent_startIngestionJobCmd.MarkFlagRequired("data-source-id")
	bedrockAgent_startIngestionJobCmd.MarkFlagRequired("knowledge-base-id")
	bedrockAgentCmd.AddCommand(bedrockAgent_startIngestionJobCmd)
}
