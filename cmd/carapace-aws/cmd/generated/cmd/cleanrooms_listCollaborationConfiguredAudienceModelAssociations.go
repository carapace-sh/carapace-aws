package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listCollaborationConfiguredAudienceModelAssociationsCmd = &cobra.Command{
	Use:   "list-collaboration-configured-audience-model-associations",
	Short: "Lists configured audience model associations within a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listCollaborationConfiguredAudienceModelAssociationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_listCollaborationConfiguredAudienceModelAssociationsCmd).Standalone()

		cleanrooms_listCollaborationConfiguredAudienceModelAssociationsCmd.Flags().String("collaboration-identifier", "", "A unique identifier for the collaboration that the configured audience model association belongs to.")
		cleanrooms_listCollaborationConfiguredAudienceModelAssociationsCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
		cleanrooms_listCollaborationConfiguredAudienceModelAssociationsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
		cleanrooms_listCollaborationConfiguredAudienceModelAssociationsCmd.MarkFlagRequired("collaboration-identifier")
	})
	cleanroomsCmd.AddCommand(cleanrooms_listCollaborationConfiguredAudienceModelAssociationsCmd)
}
