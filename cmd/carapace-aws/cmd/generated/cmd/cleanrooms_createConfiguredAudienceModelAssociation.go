package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_createConfiguredAudienceModelAssociationCmd = &cobra.Command{
	Use:   "create-configured-audience-model-association",
	Short: "Provides the details necessary to create a configured audience model association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_createConfiguredAudienceModelAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanrooms_createConfiguredAudienceModelAssociationCmd).Standalone()

		cleanrooms_createConfiguredAudienceModelAssociationCmd.Flags().String("configured-audience-model-arn", "", "A unique identifier for the configured audience model that you want to associate.")
		cleanrooms_createConfiguredAudienceModelAssociationCmd.Flags().String("configured-audience-model-association-name", "", "The name of the configured audience model association.")
		cleanrooms_createConfiguredAudienceModelAssociationCmd.Flags().String("description", "", "A description of the configured audience model association.")
		cleanrooms_createConfiguredAudienceModelAssociationCmd.Flags().Bool("manage-resource-policies", false, "When `TRUE`, indicates that the resource policy for the configured audience model resource being associated is configured for Clean Rooms to manage permissions related to the given collaboration.")
		cleanrooms_createConfiguredAudienceModelAssociationCmd.Flags().String("membership-identifier", "", "A unique identifier for one of your memberships for a collaboration.")
		cleanrooms_createConfiguredAudienceModelAssociationCmd.Flags().Bool("no-manage-resource-policies", false, "When `TRUE`, indicates that the resource policy for the configured audience model resource being associated is configured for Clean Rooms to manage permissions related to the given collaboration.")
		cleanrooms_createConfiguredAudienceModelAssociationCmd.Flags().String("tags", "", "An optional label that you can assign to a resource when you create it.")
		cleanrooms_createConfiguredAudienceModelAssociationCmd.MarkFlagRequired("configured-audience-model-arn")
		cleanrooms_createConfiguredAudienceModelAssociationCmd.MarkFlagRequired("configured-audience-model-association-name")
		cleanrooms_createConfiguredAudienceModelAssociationCmd.MarkFlagRequired("manage-resource-policies")
		cleanrooms_createConfiguredAudienceModelAssociationCmd.MarkFlagRequired("membership-identifier")
		cleanrooms_createConfiguredAudienceModelAssociationCmd.Flag("no-manage-resource-policies").Hidden = true
		cleanrooms_createConfiguredAudienceModelAssociationCmd.MarkFlagRequired("no-manage-resource-policies")
	})
	cleanroomsCmd.AddCommand(cleanrooms_createConfiguredAudienceModelAssociationCmd)
}
