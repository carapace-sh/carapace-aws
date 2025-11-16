package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_createDatasetCmd = &cobra.Command{
	Use:   "create-dataset",
	Short: "Used to create a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_createDatasetCmd).Standalone()

	iotanalytics_createDatasetCmd.Flags().String("actions", "", "A list of actions that create the dataset contents.")
	iotanalytics_createDatasetCmd.Flags().String("content-delivery-rules", "", "When dataset contents are created, they are delivered to destinations specified here.")
	iotanalytics_createDatasetCmd.Flags().String("dataset-name", "", "The name of the dataset.")
	iotanalytics_createDatasetCmd.Flags().String("late-data-rules", "", "A list of data rules that send notifications to CloudWatch, when data arrives late.")
	iotanalytics_createDatasetCmd.Flags().String("retention-period", "", "Optional.")
	iotanalytics_createDatasetCmd.Flags().String("tags", "", "Metadata which can be used to manage the dataset.")
	iotanalytics_createDatasetCmd.Flags().String("triggers", "", "A list of triggers.")
	iotanalytics_createDatasetCmd.Flags().String("versioning-configuration", "", "Optional.")
	iotanalytics_createDatasetCmd.MarkFlagRequired("actions")
	iotanalytics_createDatasetCmd.MarkFlagRequired("dataset-name")
	iotanalyticsCmd.AddCommand(iotanalytics_createDatasetCmd)
}
