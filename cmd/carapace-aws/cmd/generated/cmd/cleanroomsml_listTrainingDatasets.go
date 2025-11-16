package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listTrainingDatasetsCmd = &cobra.Command{
	Use:   "list-training-datasets",
	Short: "Returns a list of training datasets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listTrainingDatasetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_listTrainingDatasetsCmd).Standalone()

		cleanroomsml_listTrainingDatasetsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
		cleanroomsml_listTrainingDatasetsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_listTrainingDatasetsCmd)
}
