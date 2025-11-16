package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_updateDatasetCmd = &cobra.Command{
	Use:   "update-dataset",
	Short: "Updates the settings of a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_updateDatasetCmd).Standalone()

	iotanalytics_updateDatasetCmd.Flags().String("actions", "", "A list of `DatasetAction` objects.")
	iotanalytics_updateDatasetCmd.Flags().String("content-delivery-rules", "", "When dataset contents are created, they are delivered to destinations specified here.")
	iotanalytics_updateDatasetCmd.Flags().String("dataset-name", "", "The name of the dataset to update.")
	iotanalytics_updateDatasetCmd.Flags().String("late-data-rules", "", "A list of data rules that send notifications to CloudWatch, when data arrives late.")
	iotanalytics_updateDatasetCmd.Flags().String("retention-period", "", "How long, in days, dataset contents are kept for the dataset.")
	iotanalytics_updateDatasetCmd.Flags().String("triggers", "", "A list of `DatasetTrigger` objects.")
	iotanalytics_updateDatasetCmd.Flags().String("versioning-configuration", "", "Optional.")
	iotanalytics_updateDatasetCmd.MarkFlagRequired("actions")
	iotanalytics_updateDatasetCmd.MarkFlagRequired("dataset-name")
	iotanalyticsCmd.AddCommand(iotanalytics_updateDatasetCmd)
}
