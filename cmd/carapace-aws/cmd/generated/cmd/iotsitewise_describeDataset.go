package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeDatasetCmd = &cobra.Command{
	Use:   "describe-dataset",
	Short: "Retrieves information about a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_describeDatasetCmd).Standalone()

		iotsitewise_describeDatasetCmd.Flags().String("dataset-id", "", "The ID of the dataset.")
		iotsitewise_describeDatasetCmd.MarkFlagRequired("dataset-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_describeDatasetCmd)
}
