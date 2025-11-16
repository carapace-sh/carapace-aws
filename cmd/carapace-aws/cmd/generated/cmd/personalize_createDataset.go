package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createDatasetCmd = &cobra.Command{
	Use:   "create-dataset",
	Short: "Creates an empty dataset and adds it to the specified dataset group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createDatasetCmd).Standalone()

	personalize_createDatasetCmd.Flags().String("dataset-group-arn", "", "The Amazon Resource Name (ARN) of the dataset group to add the dataset to.")
	personalize_createDatasetCmd.Flags().String("dataset-type", "", "The type of dataset.")
	personalize_createDatasetCmd.Flags().String("name", "", "The name for the dataset.")
	personalize_createDatasetCmd.Flags().String("schema-arn", "", "The ARN of the schema to associate with the dataset.")
	personalize_createDatasetCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html) to apply to the dataset.")
	personalize_createDatasetCmd.MarkFlagRequired("dataset-group-arn")
	personalize_createDatasetCmd.MarkFlagRequired("dataset-type")
	personalize_createDatasetCmd.MarkFlagRequired("name")
	personalize_createDatasetCmd.MarkFlagRequired("schema-arn")
	personalizeCmd.AddCommand(personalize_createDatasetCmd)
}
