package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_deleteConfiguredAudienceModelCmd = &cobra.Command{
	Use:   "delete-configured-audience-model",
	Short: "Deletes the specified configured audience model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_deleteConfiguredAudienceModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_deleteConfiguredAudienceModelCmd).Standalone()

		cleanroomsml_deleteConfiguredAudienceModelCmd.Flags().String("configured-audience-model-arn", "", "The Amazon Resource Name (ARN) of the configured audience model that you want to delete.")
		cleanroomsml_deleteConfiguredAudienceModelCmd.MarkFlagRequired("configured-audience-model-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_deleteConfiguredAudienceModelCmd)
}
