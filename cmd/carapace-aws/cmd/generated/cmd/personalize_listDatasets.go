package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listDatasetsCmd = &cobra.Command{
	Use:   "list-datasets",
	Short: "Returns the list of datasets contained in the given dataset group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listDatasetsCmd).Standalone()

	personalize_listDatasetsCmd.Flags().String("dataset-group-arn", "", "The Amazon Resource Name (ARN) of the dataset group that contains the datasets to list.")
	personalize_listDatasetsCmd.Flags().String("max-results", "", "The maximum number of datasets to return.")
	personalize_listDatasetsCmd.Flags().String("next-token", "", "A token returned from the previous call to `ListDatasets` for getting the next set of dataset import jobs (if they exist).")
	personalizeCmd.AddCommand(personalize_listDatasetsCmd)
}
