package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_updateConfiguredAudienceModelAssociationCmd = &cobra.Command{
	Use:   "update-configured-audience-model-association",
	Short: "Provides the details necessary to update a configured audience model association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_updateConfiguredAudienceModelAssociationCmd).Standalone()

	cleanrooms_updateConfiguredAudienceModelAssociationCmd.Flags().String("configured-audience-model-association-identifier", "", "A unique identifier for the configured audience model association that you want to update.")
	cleanrooms_updateConfiguredAudienceModelAssociationCmd.Flags().String("description", "", "A new description for the configured audience model association.")
	cleanrooms_updateConfiguredAudienceModelAssociationCmd.Flags().String("membership-identifier", "", "A unique identifier of the membership that contains the configured audience model association that you want to update.")
	cleanrooms_updateConfiguredAudienceModelAssociationCmd.Flags().String("name", "", "A new name for the configured audience model association.")
	cleanrooms_updateConfiguredAudienceModelAssociationCmd.MarkFlagRequired("configured-audience-model-association-identifier")
	cleanrooms_updateConfiguredAudienceModelAssociationCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_updateConfiguredAudienceModelAssociationCmd)
}
