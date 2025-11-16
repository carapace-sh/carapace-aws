package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_deleteDatasetCmd = &cobra.Command{
	Use:   "delete-dataset",
	Short: "Deletes a dataset from DataBrew.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_deleteDatasetCmd).Standalone()

	databrew_deleteDatasetCmd.Flags().String("name", "", "The name of the dataset to be deleted.")
	databrew_deleteDatasetCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_deleteDatasetCmd)
}
