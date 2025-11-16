package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_getIngestionJobCmd = &cobra.Command{
	Use:   "get-ingestion-job",
	Short: "Gets information about a data ingestion job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_getIngestionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_getIngestionJobCmd).Standalone()

		bedrockAgent_getIngestionJobCmd.Flags().String("data-source-id", "", "The unique identifier of the data source for the data ingestion job you want to get information on.")
		bedrockAgent_getIngestionJobCmd.Flags().String("ingestion-job-id", "", "The unique identifier of the data ingestion job you want to get information on.")
		bedrockAgent_getIngestionJobCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base for the data ingestion job you want to get information on.")
		bedrockAgent_getIngestionJobCmd.MarkFlagRequired("data-source-id")
		bedrockAgent_getIngestionJobCmd.MarkFlagRequired("ingestion-job-id")
		bedrockAgent_getIngestionJobCmd.MarkFlagRequired("knowledge-base-id")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_getIngestionJobCmd)
}
