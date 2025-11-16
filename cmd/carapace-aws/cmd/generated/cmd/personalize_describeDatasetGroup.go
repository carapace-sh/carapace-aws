package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeDatasetGroupCmd = &cobra.Command{
	Use:   "describe-dataset-group",
	Short: "Describes the given dataset group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeDatasetGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_describeDatasetGroupCmd).Standalone()

		personalize_describeDatasetGroupCmd.Flags().String("dataset-group-arn", "", "The Amazon Resource Name (ARN) of the dataset group to describe.")
		personalize_describeDatasetGroupCmd.MarkFlagRequired("dataset-group-arn")
	})
	personalizeCmd.AddCommand(personalize_describeDatasetGroupCmd)
}
