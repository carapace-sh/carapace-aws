package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_createDataSourceCmd = &cobra.Command{
	Use:   "create-data-source",
	Short: "Connects a knowledge base to a data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_createDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrockAgent_createDataSourceCmd).Standalone()

		bedrockAgent_createDataSourceCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrockAgent_createDataSourceCmd.Flags().String("data-deletion-policy", "", "The data deletion policy for the data source.")
		bedrockAgent_createDataSourceCmd.Flags().String("data-source-configuration", "", "The connection configuration for the data source.")
		bedrockAgent_createDataSourceCmd.Flags().String("description", "", "A description of the data source.")
		bedrockAgent_createDataSourceCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base to which to add the data source.")
		bedrockAgent_createDataSourceCmd.Flags().String("name", "", "The name of the data source.")
		bedrockAgent_createDataSourceCmd.Flags().String("server-side-encryption-configuration", "", "Contains details about the server-side encryption for the data source.")
		bedrockAgent_createDataSourceCmd.Flags().String("vector-ingestion-configuration", "", "Contains details about how to ingest the documents in the data source.")
		bedrockAgent_createDataSourceCmd.MarkFlagRequired("data-source-configuration")
		bedrockAgent_createDataSourceCmd.MarkFlagRequired("knowledge-base-id")
		bedrockAgent_createDataSourceCmd.MarkFlagRequired("name")
	})
	bedrockAgentCmd.AddCommand(bedrockAgent_createDataSourceCmd)
}
