package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeDatasetCmd = &cobra.Command{
	Use:   "describe-dataset",
	Short: "Describes the given dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_describeDatasetCmd).Standalone()

		personalize_describeDatasetCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset to describe.")
		personalize_describeDatasetCmd.MarkFlagRequired("dataset-arn")
	})
	personalizeCmd.AddCommand(personalize_describeDatasetCmd)
}
