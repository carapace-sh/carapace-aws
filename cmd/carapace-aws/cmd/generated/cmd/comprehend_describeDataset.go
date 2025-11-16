package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeDatasetCmd = &cobra.Command{
	Use:   "describe-dataset",
	Short: "Returns information about the dataset that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeDatasetCmd).Standalone()

	comprehend_describeDatasetCmd.Flags().String("dataset-arn", "", "The ARN of the dataset.")
	comprehend_describeDatasetCmd.MarkFlagRequired("dataset-arn")
	comprehendCmd.AddCommand(comprehend_describeDatasetCmd)
}
