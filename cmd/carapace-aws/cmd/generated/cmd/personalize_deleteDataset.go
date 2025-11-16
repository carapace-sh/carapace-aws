package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_deleteDatasetCmd = &cobra.Command{
	Use:   "delete-dataset",
	Short: "Deletes a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_deleteDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_deleteDatasetCmd).Standalone()

		personalize_deleteDatasetCmd.Flags().String("dataset-arn", "", "The Amazon Resource Name (ARN) of the dataset to delete.")
		personalize_deleteDatasetCmd.MarkFlagRequired("dataset-arn")
	})
	personalizeCmd.AddCommand(personalize_deleteDatasetCmd)
}
