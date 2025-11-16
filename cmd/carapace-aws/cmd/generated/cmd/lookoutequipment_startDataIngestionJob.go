package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_startDataIngestionJobCmd = &cobra.Command{
	Use:   "start-data-ingestion-job",
	Short: "Starts a data ingestion job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_startDataIngestionJobCmd).Standalone()

	lookoutequipment_startDataIngestionJobCmd.Flags().String("client-token", "", "A unique identifier for the request.")
	lookoutequipment_startDataIngestionJobCmd.Flags().String("dataset-name", "", "The name of the dataset being used by the data ingestion job.")
	lookoutequipment_startDataIngestionJobCmd.Flags().String("ingestion-input-configuration", "", "Specifies information for the input data for the data ingestion job, including dataset S3 location.")
	lookoutequipment_startDataIngestionJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of a role with permission to access the data source for the data ingestion job.")
	lookoutequipment_startDataIngestionJobCmd.MarkFlagRequired("client-token")
	lookoutequipment_startDataIngestionJobCmd.MarkFlagRequired("dataset-name")
	lookoutequipment_startDataIngestionJobCmd.MarkFlagRequired("ingestion-input-configuration")
	lookoutequipment_startDataIngestionJobCmd.MarkFlagRequired("role-arn")
	lookoutequipmentCmd.AddCommand(lookoutequipment_startDataIngestionJobCmd)
}
