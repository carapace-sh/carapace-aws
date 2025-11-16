package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getAudienceModelCmd = &cobra.Command{
	Use:   "get-audience-model",
	Short: "Returns information about an audience model",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getAudienceModelCmd).Standalone()

	cleanroomsml_getAudienceModelCmd.Flags().String("audience-model-arn", "", "The Amazon Resource Name (ARN) of the audience model that you are interested in.")
	cleanroomsml_getAudienceModelCmd.MarkFlagRequired("audience-model-arn")
	cleanroomsmlCmd.AddCommand(cleanroomsml_getAudienceModelCmd)
}
