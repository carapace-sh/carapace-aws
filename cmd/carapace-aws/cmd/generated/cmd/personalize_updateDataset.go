package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_updateDatasetCmd = &cobra.Command{
	Use:   "update-dataset",
	Short: "Update a dataset to replace its schema with a new or existing one.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_updateDatasetCmd).Standalone()

	personalize_updateDatasetCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset that you want to update.")
	personalize_updateDatasetCmd.Flags().String("schema-arn", "", "The Amazon Resource Name (ARN) of the new schema you want use.")
	personalize_updateDatasetCmd.MarkFlagRequired("dataset-arn")
	personalize_updateDatasetCmd.MarkFlagRequired("schema-arn")
	personalizeCmd.AddCommand(personalize_updateDatasetCmd)
}
