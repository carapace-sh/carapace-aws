package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getConfiguredAudienceModelCmd = &cobra.Command{
	Use:   "get-configured-audience-model",
	Short: "Returns information about a specified configured audience model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getConfiguredAudienceModelCmd).Standalone()

	cleanroomsml_getConfiguredAudienceModelCmd.Flags().String("configured-audience-model-arn", "", "The Amazon Resource Name (ARN) of the configured audience model that you are interested in.")
	cleanroomsml_getConfiguredAudienceModelCmd.MarkFlagRequired("configured-audience-model-arn")
	cleanroomsmlCmd.AddCommand(cleanroomsml_getConfiguredAudienceModelCmd)
}
