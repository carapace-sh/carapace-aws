package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getConfiguredAudienceModelAssociationCmd = &cobra.Command{
	Use:   "get-configured-audience-model-association",
	Short: "Returns information about a configured audience model association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getConfiguredAudienceModelAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_getConfiguredAudienceModelAssociationCmd).Standalone()

		cleanrooms_getConfiguredAudienceModelAssociationCmd.Flags().String("configured-audience-model-association-identifier", "", "A unique identifier for the configured audience model association that you want to retrieve.")
		cleanrooms_getConfiguredAudienceModelAssociationCmd.Flags().String("membership-identifier", "", "A unique identifier for the membership that contains the configured audience model association that you want to retrieve.")
		cleanrooms_getConfiguredAudienceModelAssociationCmd.MarkFlagRequired("configured-audience-model-association-identifier")
		cleanrooms_getConfiguredAudienceModelAssociationCmd.MarkFlagRequired("membership-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_getConfiguredAudienceModelAssociationCmd)
}
