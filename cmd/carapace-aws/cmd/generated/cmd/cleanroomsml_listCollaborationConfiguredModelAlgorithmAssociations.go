package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listCollaborationConfiguredModelAlgorithmAssociationsCmd = &cobra.Command{
	Use:   "list-collaboration-configured-model-algorithm-associations",
	Short: "Returns a list of the configured model algorithm associations in a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listCollaborationConfiguredModelAlgorithmAssociationsCmd).Standalone()

	cleanroomsml_listCollaborationConfiguredModelAlgorithmAssociationsCmd.Flags().String("collaboration-identifier", "", "The collaboration ID of the collaboration that contains the configured model algorithm associations that you are interested in.")
	cleanroomsml_listCollaborationConfiguredModelAlgorithmAssociationsCmd.Flags().String("max-results", "", "The maximum size of the results that is returned per call.")
	cleanroomsml_listCollaborationConfiguredModelAlgorithmAssociationsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	cleanroomsml_listCollaborationConfiguredModelAlgorithmAssociationsCmd.MarkFlagRequired("collaboration-identifier")
	cleanroomsmlCmd.AddCommand(cleanroomsml_listCollaborationConfiguredModelAlgorithmAssociationsCmd)
}
