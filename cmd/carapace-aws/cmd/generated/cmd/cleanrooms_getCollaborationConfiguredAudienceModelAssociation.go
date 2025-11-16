package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_getCollaborationConfiguredAudienceModelAssociationCmd = &cobra.Command{
	Use:   "get-collaboration-configured-audience-model-association",
	Short: "Retrieves a configured audience model association within a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_getCollaborationConfiguredAudienceModelAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_getCollaborationConfiguredAudienceModelAssociationCmd).Standalone()

		cleanrooms_getCollaborationConfiguredAudienceModelAssociationCmd.Flags().String("collaboration-identifier", "", "A unique identifier for the collaboration that the configured audience model association belongs to.")
		cleanrooms_getCollaborationConfiguredAudienceModelAssociationCmd.Flags().String("configured-audience-model-association-identifier", "", "A unique identifier for the configured audience model association that you want to retrieve.")
		cleanrooms_getCollaborationConfiguredAudienceModelAssociationCmd.MarkFlagRequired("collaboration-identifier")
		cleanrooms_getCollaborationConfiguredAudienceModelAssociationCmd.MarkFlagRequired("configured-audience-model-association-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_getCollaborationConfiguredAudienceModelAssociationCmd)
}
