package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_createConfiguredModelAlgorithmAssociationCmd = &cobra.Command{
	Use:   "create-configured-model-algorithm-association",
	Short: "Associates a configured model algorithm to a collaboration for use by any member of the collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_createConfiguredModelAlgorithmAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_createConfiguredModelAlgorithmAssociationCmd).Standalone()

		cleanroomsml_createConfiguredModelAlgorithmAssociationCmd.Flags().String("configured-model-algorithm-arn", "", "The Amazon Resource Name (ARN) of the configured model algorithm that you want to associate.")
		cleanroomsml_createConfiguredModelAlgorithmAssociationCmd.Flags().String("description", "", "The description of the configured model algorithm association.")
		cleanroomsml_createConfiguredModelAlgorithmAssociationCmd.Flags().String("membership-identifier", "", "The membership ID of the member who is associating this configured model algorithm.")
		cleanroomsml_createConfiguredModelAlgorithmAssociationCmd.Flags().String("name", "", "The name of the configured model algorithm association.")
		cleanroomsml_createConfiguredModelAlgorithmAssociationCmd.Flags().String("privacy-configuration", "", "Specifies the privacy configuration information for the configured model algorithm association.")
		cleanroomsml_createConfiguredModelAlgorithmAssociationCmd.Flags().String("tags", "", "The optional metadata that you apply to the resource to help you categorize and organize them.")
		cleanroomsml_createConfiguredModelAlgorithmAssociationCmd.MarkFlagRequired("configured-model-algorithm-arn")
		cleanroomsml_createConfiguredModelAlgorithmAssociationCmd.MarkFlagRequired("membership-identifier")
		cleanroomsml_createConfiguredModelAlgorithmAssociationCmd.MarkFlagRequired("name")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_createConfiguredModelAlgorithmAssociationCmd)
}
