package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_deleteConfiguredModelAlgorithmCmd = &cobra.Command{
	Use:   "delete-configured-model-algorithm",
	Short: "Deletes a configured model algorithm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_deleteConfiguredModelAlgorithmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_deleteConfiguredModelAlgorithmCmd).Standalone()

		cleanroomsml_deleteConfiguredModelAlgorithmCmd.Flags().String("configured-model-algorithm-arn", "", "The Amazon Resource Name (ARN) of the configured model algorithm that you want to delete.")
		cleanroomsml_deleteConfiguredModelAlgorithmCmd.MarkFlagRequired("configured-model-algorithm-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_deleteConfiguredModelAlgorithmCmd)
}
