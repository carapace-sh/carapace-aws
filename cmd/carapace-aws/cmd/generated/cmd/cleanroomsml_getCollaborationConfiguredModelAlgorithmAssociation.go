package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getCollaborationConfiguredModelAlgorithmAssociationCmd = &cobra.Command{
	Use:   "get-collaboration-configured-model-algorithm-association",
	Short: "Returns information about the configured model algorithm association in a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getCollaborationConfiguredModelAlgorithmAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_getCollaborationConfiguredModelAlgorithmAssociationCmd).Standalone()

		cleanroomsml_getCollaborationConfiguredModelAlgorithmAssociationCmd.Flags().String("collaboration-identifier", "", "The collaboration ID for the collaboration that contains the configured model algorithm association that you want to return information about.")
		cleanroomsml_getCollaborationConfiguredModelAlgorithmAssociationCmd.Flags().String("configured-model-algorithm-association-arn", "", "The Amazon Resource Name (ARN) of the configured model algorithm association that you want to return information about.")
		cleanroomsml_getCollaborationConfiguredModelAlgorithmAssociationCmd.MarkFlagRequired("collaboration-identifier")
		cleanroomsml_getCollaborationConfiguredModelAlgorithmAssociationCmd.MarkFlagRequired("configured-model-algorithm-association-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_getCollaborationConfiguredModelAlgorithmAssociationCmd)
}
