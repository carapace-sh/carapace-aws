package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_deleteDatasetGroupCmd = &cobra.Command{
	Use:   "delete-dataset-group",
	Short: "Deletes a dataset group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_deleteDatasetGroupCmd).Standalone()

	personalize_deleteDatasetGroupCmd.Flags().String("dataset-group-arn", "", "The ARN of the dataset group to delete.")
	personalize_deleteDatasetGroupCmd.MarkFlagRequired("dataset-group-arn")
	personalizeCmd.AddCommand(personalize_deleteDatasetGroupCmd)
}
