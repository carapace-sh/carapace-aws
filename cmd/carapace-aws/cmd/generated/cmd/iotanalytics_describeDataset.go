package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotanalytics_describeDatasetCmd = &cobra.Command{
	Use:   "describe-dataset",
	Short: "Retrieves information about a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotanalytics_describeDatasetCmd).Standalone()

	iotanalytics_describeDatasetCmd.Flags().String("dataset-name", "", "The name of the dataset whose information is retrieved.")
	iotanalytics_describeDatasetCmd.MarkFlagRequired("dataset-name")
	iotanalyticsCmd.AddCommand(iotanalytics_describeDatasetCmd)
}
