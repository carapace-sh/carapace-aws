package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_describeDataIngestionJobCmd = &cobra.Command{
	Use:   "describe-data-ingestion-job",
	Short: "Provides information on a specific data ingestion job such as creation time, dataset ARN, and status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_describeDataIngestionJobCmd).Standalone()

	lookoutequipment_describeDataIngestionJobCmd.Flags().String("job-id", "", "The job ID of the data ingestion job.")
	lookoutequipment_describeDataIngestionJobCmd.MarkFlagRequired("job-id")
	lookoutequipmentCmd.AddCommand(lookoutequipment_describeDataIngestionJobCmd)
}
