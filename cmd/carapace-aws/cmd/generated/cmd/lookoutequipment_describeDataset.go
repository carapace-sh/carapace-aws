package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_describeDatasetCmd = &cobra.Command{
	Use:   "describe-dataset",
	Short: "Provides a JSON description of the data in each time series dataset, including names, column names, and data types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_describeDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_describeDatasetCmd).Standalone()

		lookoutequipment_describeDatasetCmd.Flags().String("dataset-name", "", "The name of the dataset to be described.")
		lookoutequipment_describeDatasetCmd.MarkFlagRequired("dataset-name")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_describeDatasetCmd)
}
