package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_createAudienceModelCmd = &cobra.Command{
	Use:   "create-audience-model",
	Short: "Defines the information necessary to create an audience model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_createAudienceModelCmd).Standalone()

	cleanroomsml_createAudienceModelCmd.Flags().String("description", "", "The description of the audience model.")
	cleanroomsml_createAudienceModelCmd.Flags().String("kms-key-arn", "", "The Amazon Resource Name (ARN) of the KMS key.")
	cleanroomsml_createAudienceModelCmd.Flags().String("name", "", "The name of the audience model resource.")
	cleanroomsml_createAudienceModelCmd.Flags().String("tags", "", "The optional metadata that you apply to the resource to help you categorize and organize them.")
	cleanroomsml_createAudienceModelCmd.Flags().String("training-data-end-time", "", "The end date and time of the training window.")
	cleanroomsml_createAudienceModelCmd.Flags().String("training-data-start-time", "", "The start date and time of the training window.")
	cleanroomsml_createAudienceModelCmd.Flags().String("training-dataset-arn", "", "The Amazon Resource Name (ARN) of the training dataset for this audience model.")
	cleanroomsml_createAudienceModelCmd.MarkFlagRequired("name")
	cleanroomsml_createAudienceModelCmd.MarkFlagRequired("training-dataset-arn")
	cleanroomsmlCmd.AddCommand(cleanroomsml_createAudienceModelCmd)
}
