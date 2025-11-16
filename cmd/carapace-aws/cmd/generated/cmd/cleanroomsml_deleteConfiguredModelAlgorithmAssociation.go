package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_deleteConfiguredModelAlgorithmAssociationCmd = &cobra.Command{
	Use:   "delete-configured-model-algorithm-association",
	Short: "Deletes a configured model algorithm association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_deleteConfiguredModelAlgorithmAssociationCmd).Standalone()

	cleanroomsml_deleteConfiguredModelAlgorithmAssociationCmd.Flags().String("configured-model-algorithm-association-arn", "", "The Amazon Resource Name (ARN) of the configured model algorithm association that you want to delete.")
	cleanroomsml_deleteConfiguredModelAlgorithmAssociationCmd.Flags().String("membership-identifier", "", "The membership ID of the member that is deleting the configured model algorithm association.")
	cleanroomsml_deleteConfiguredModelAlgorithmAssociationCmd.MarkFlagRequired("configured-model-algorithm-association-arn")
	cleanroomsml_deleteConfiguredModelAlgorithmAssociationCmd.MarkFlagRequired("membership-identifier")
	cleanroomsmlCmd.AddCommand(cleanroomsml_deleteConfiguredModelAlgorithmAssociationCmd)
}
