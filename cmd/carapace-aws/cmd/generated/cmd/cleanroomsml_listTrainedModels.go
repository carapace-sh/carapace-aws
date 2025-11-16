package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listTrainedModelsCmd = &cobra.Command{
	Use:   "list-trained-models",
	Short: "Returns a list of trained models.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listTrainedModelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_listTrainedModelsCmd).Standalone()

		cleanroomsml_listTrainedModelsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
		cleanroomsml_listTrainedModelsCmd.Flags().String("membership-identifier", "", "The membership ID of the member that created the trained models you are interested in.")
		cleanroomsml_listTrainedModelsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
		cleanroomsml_listTrainedModelsCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_listTrainedModelsCmd)
}
