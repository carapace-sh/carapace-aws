package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listTrainedModelVersionsCmd = &cobra.Command{
	Use:   "list-trained-model-versions",
	Short: "Returns a list of trained model versions for a specified trained model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listTrainedModelVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_listTrainedModelVersionsCmd).Standalone()

		cleanroomsml_listTrainedModelVersionsCmd.Flags().String("max-results", "", "The maximum number of trained model versions to return in a single page.")
		cleanroomsml_listTrainedModelVersionsCmd.Flags().String("membership-identifier", "", "The membership identifier for the collaboration that contains the trained model.")
		cleanroomsml_listTrainedModelVersionsCmd.Flags().String("next-token", "", "The pagination token from a previous `ListTrainedModelVersions` request.")
		cleanroomsml_listTrainedModelVersionsCmd.Flags().String("status", "", "Filter the results to only include trained model versions with the specified status.")
		cleanroomsml_listTrainedModelVersionsCmd.Flags().String("trained-model-arn", "", "The Amazon Resource Name (ARN) of the trained model for which to list versions.")
		cleanroomsml_listTrainedModelVersionsCmd.MarkFlagRequired("membership-identifier")
		cleanroomsml_listTrainedModelVersionsCmd.MarkFlagRequired("trained-model-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_listTrainedModelVersionsCmd)
}
