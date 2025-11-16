package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_deleteAudienceModelCmd = &cobra.Command{
	Use:   "delete-audience-model",
	Short: "Specifies an audience model that you want to delete.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_deleteAudienceModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_deleteAudienceModelCmd).Standalone()

		cleanroomsml_deleteAudienceModelCmd.Flags().String("audience-model-arn", "", "The Amazon Resource Name (ARN) of the audience model that you want to delete.")
		cleanroomsml_deleteAudienceModelCmd.MarkFlagRequired("audience-model-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_deleteAudienceModelCmd)
}
