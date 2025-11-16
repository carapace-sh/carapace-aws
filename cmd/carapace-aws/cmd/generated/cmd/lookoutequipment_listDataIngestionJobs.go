package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_listDataIngestionJobsCmd = &cobra.Command{
	Use:   "list-data-ingestion-jobs",
	Short: "Provides a list of all data ingestion jobs, including dataset name and ARN, S3 location of the input data, status, and so on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_listDataIngestionJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_listDataIngestionJobsCmd).Standalone()

		lookoutequipment_listDataIngestionJobsCmd.Flags().String("dataset-name", "", "The name of the dataset being used for the data ingestion job.")
		lookoutequipment_listDataIngestionJobsCmd.Flags().String("max-results", "", "Specifies the maximum number of data ingestion jobs to list.")
		lookoutequipment_listDataIngestionJobsCmd.Flags().String("next-token", "", "An opaque pagination token indicating where to continue the listing of data ingestion jobs.")
		lookoutequipment_listDataIngestionJobsCmd.Flags().String("status", "", "Indicates the status of the data ingestion job.")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_listDataIngestionJobsCmd)
}
