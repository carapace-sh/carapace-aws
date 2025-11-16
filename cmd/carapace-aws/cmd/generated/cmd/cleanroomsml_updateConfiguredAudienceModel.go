package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_updateConfiguredAudienceModelCmd = &cobra.Command{
	Use:   "update-configured-audience-model",
	Short: "Provides the information necessary to update a configured audience model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_updateConfiguredAudienceModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_updateConfiguredAudienceModelCmd).Standalone()

		cleanroomsml_updateConfiguredAudienceModelCmd.Flags().String("audience-model-arn", "", "The Amazon Resource Name (ARN) of the new audience model that you want to use.")
		cleanroomsml_updateConfiguredAudienceModelCmd.Flags().String("audience-size-config", "", "The new audience size configuration.")
		cleanroomsml_updateConfiguredAudienceModelCmd.Flags().String("configured-audience-model-arn", "", "The Amazon Resource Name (ARN) of the configured audience model that you want to update.")
		cleanroomsml_updateConfiguredAudienceModelCmd.Flags().String("description", "", "The new description of the configured audience model.")
		cleanroomsml_updateConfiguredAudienceModelCmd.Flags().String("min-matching-seed-size", "", "The minimum number of users from the seed audience that must match with users in the training data of the audience model.")
		cleanroomsml_updateConfiguredAudienceModelCmd.Flags().String("output-config", "", "The new output configuration.")
		cleanroomsml_updateConfiguredAudienceModelCmd.Flags().String("shared-audience-metrics", "", "The new value for whether to share audience metrics.")
		cleanroomsml_updateConfiguredAudienceModelCmd.MarkFlagRequired("configured-audience-model-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_updateConfiguredAudienceModelCmd)
}
