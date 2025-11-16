package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrockAgent_updateDataSourceCmd = &cobra.Command{
	Use:   "update-data-source",
	Short: "Updates the configurations for a data source connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrockAgent_updateDataSourceCmd).Standalone()

	bedrockAgent_updateDataSourceCmd.Flags().String("data-deletion-policy", "", "The data deletion policy for the data source that you want to update.")
	bedrockAgent_updateDataSourceCmd.Flags().String("data-source-configuration", "", "The connection configuration for the data source that you want to update.")
	bedrockAgent_updateDataSourceCmd.Flags().String("data-source-id", "", "The unique identifier of the data source.")
	bedrockAgent_updateDataSourceCmd.Flags().String("description", "", "Specifies a new description for the data source.")
	bedrockAgent_updateDataSourceCmd.Flags().String("knowledge-base-id", "", "The unique identifier of the knowledge base for the data source.")
	bedrockAgent_updateDataSourceCmd.Flags().String("name", "", "Specifies a new name for the data source.")
	bedrockAgent_updateDataSourceCmd.Flags().String("server-side-encryption-configuration", "", "Contains details about server-side encryption of the data source.")
	bedrockAgent_updateDataSourceCmd.Flags().String("vector-ingestion-configuration", "", "Contains details about how to ingest the documents in the data source.")
	bedrockAgent_updateDataSourceCmd.MarkFlagRequired("data-source-configuration")
	bedrockAgent_updateDataSourceCmd.MarkFlagRequired("data-source-id")
	bedrockAgent_updateDataSourceCmd.MarkFlagRequired("knowledge-base-id")
	bedrockAgent_updateDataSourceCmd.MarkFlagRequired("name")
	bedrockAgentCmd.AddCommand(bedrockAgent_updateDataSourceCmd)
}
