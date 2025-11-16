package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_deleteConfiguredAudienceModelAssociationCmd = &cobra.Command{
	Use:   "delete-configured-audience-model-association",
	Short: "Provides the information necessary to delete a configured audience model association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_deleteConfiguredAudienceModelAssociationCmd).Standalone()

	cleanrooms_deleteConfiguredAudienceModelAssociationCmd.Flags().String("configured-audience-model-association-identifier", "", "A unique identifier of the configured audience model association that you want to delete.")
	cleanrooms_deleteConfiguredAudienceModelAssociationCmd.Flags().String("membership-identifier", "", "A unique identifier of the membership that contains the audience model association that you want to delete.")
	cleanrooms_deleteConfiguredAudienceModelAssociationCmd.MarkFlagRequired("configured-audience-model-association-identifier")
	cleanrooms_deleteConfiguredAudienceModelAssociationCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_deleteConfiguredAudienceModelAssociationCmd)
}
