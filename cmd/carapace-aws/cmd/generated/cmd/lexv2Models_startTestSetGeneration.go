package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lexv2Models_startTestSetGenerationCmd = &cobra.Command{
	Use:   "start-test-set-generation",
	Short: "The action to start the generation of test set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lexv2Models_startTestSetGenerationCmd).Standalone()

	lexv2Models_startTestSetGenerationCmd.Flags().String("description", "", "The test set description for the test set generation request.")
	lexv2Models_startTestSetGenerationCmd.Flags().String("generation-data-source", "", "The data source for the test set generation.")
	lexv2Models_startTestSetGenerationCmd.Flags().String("role-arn", "", "The roleARN used for any operation in the test set to access resources in the Amazon Web Services account.")
	lexv2Models_startTestSetGenerationCmd.Flags().String("storage-location", "", "The Amazon S3 storage location for the test set generation.")
	lexv2Models_startTestSetGenerationCmd.Flags().String("test-set-name", "", "The test set name for the test set generation request.")
	lexv2Models_startTestSetGenerationCmd.Flags().String("test-set-tags", "", "A list of tags to add to the test set.")
	lexv2Models_startTestSetGenerationCmd.MarkFlagRequired("generation-data-source")
	lexv2Models_startTestSetGenerationCmd.MarkFlagRequired("role-arn")
	lexv2Models_startTestSetGenerationCmd.MarkFlagRequired("storage-location")
	lexv2Models_startTestSetGenerationCmd.MarkFlagRequired("test-set-name")
	lexv2ModelsCmd.AddCommand(lexv2Models_startTestSetGenerationCmd)
}
