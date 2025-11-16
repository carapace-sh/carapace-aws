package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getConfiguredModelAlgorithmAssociationCmd = &cobra.Command{
	Use:   "get-configured-model-algorithm-association",
	Short: "Returns information about a configured model algorithm association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getConfiguredModelAlgorithmAssociationCmd).Standalone()

	cleanroomsml_getConfiguredModelAlgorithmAssociationCmd.Flags().String("configured-model-algorithm-association-arn", "", "The Amazon Resource Name (ARN) of the configured model algorithm association that you want to return information about.")
	cleanroomsml_getConfiguredModelAlgorithmAssociationCmd.Flags().String("membership-identifier", "", "The membership ID of the member that created the configured model algorithm association.")
	cleanroomsml_getConfiguredModelAlgorithmAssociationCmd.MarkFlagRequired("configured-model-algorithm-association-arn")
	cleanroomsml_getConfiguredModelAlgorithmAssociationCmd.MarkFlagRequired("membership-identifier")
	cleanroomsmlCmd.AddCommand(cleanroomsml_getConfiguredModelAlgorithmAssociationCmd)
}
