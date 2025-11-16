package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_createConfiguredAudienceModelCmd = &cobra.Command{
	Use:   "create-configured-audience-model",
	Short: "Defines the information necessary to create a configured audience model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_createConfiguredAudienceModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_createConfiguredAudienceModelCmd).Standalone()

		cleanroomsml_createConfiguredAudienceModelCmd.Flags().String("audience-model-arn", "", "The Amazon Resource Name (ARN) of the audience model to use for the configured audience model.")
		cleanroomsml_createConfiguredAudienceModelCmd.Flags().String("audience-size-config", "", "Configure the list of output sizes of audiences that can be created using this configured audience model.")
		cleanroomsml_createConfiguredAudienceModelCmd.Flags().String("child-resource-tag-on-create-policy", "", "Configure how the service tags audience generation jobs created using this configured audience model.")
		cleanroomsml_createConfiguredAudienceModelCmd.Flags().String("description", "", "The description of the configured audience model.")
		cleanroomsml_createConfiguredAudienceModelCmd.Flags().String("min-matching-seed-size", "", "The minimum number of users from the seed audience that must match with users in the training data of the audience model.")
		cleanroomsml_createConfiguredAudienceModelCmd.Flags().String("name", "", "The name of the configured audience model.")
		cleanroomsml_createConfiguredAudienceModelCmd.Flags().String("output-config", "", "Configure the Amazon S3 location and IAM Role for audiences created using this configured audience model.")
		cleanroomsml_createConfiguredAudienceModelCmd.Flags().String("shared-audience-metrics", "", "Whether audience metrics are shared.")
		cleanroomsml_createConfiguredAudienceModelCmd.Flags().String("tags", "", "The optional metadata that you apply to the resource to help you categorize and organize them.")
		cleanroomsml_createConfiguredAudienceModelCmd.MarkFlagRequired("audience-model-arn")
		cleanroomsml_createConfiguredAudienceModelCmd.MarkFlagRequired("name")
		cleanroomsml_createConfiguredAudienceModelCmd.MarkFlagRequired("output-config")
		cleanroomsml_createConfiguredAudienceModelCmd.MarkFlagRequired("shared-audience-metrics")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_createConfiguredAudienceModelCmd)
}
