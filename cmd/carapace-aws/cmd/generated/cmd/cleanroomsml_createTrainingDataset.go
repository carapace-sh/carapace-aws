package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_createTrainingDatasetCmd = &cobra.Command{
	Use:   "create-training-dataset",
	Short: "Defines the information necessary to create a training dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_createTrainingDatasetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_createTrainingDatasetCmd).Standalone()

		cleanroomsml_createTrainingDatasetCmd.Flags().String("description", "", "The description of the training dataset.")
		cleanroomsml_createTrainingDatasetCmd.Flags().String("name", "", "The name of the training dataset.")
		cleanroomsml_createTrainingDatasetCmd.Flags().String("role-arn", "", "The ARN of the IAM role that Clean Rooms ML can assume to read the data referred to in the `dataSource` field of each dataset.")
		cleanroomsml_createTrainingDatasetCmd.Flags().String("tags", "", "The optional metadata that you apply to the resource to help you categorize and organize them.")
		cleanroomsml_createTrainingDatasetCmd.Flags().String("training-data", "", "An array of information that lists the Dataset objects, which specifies the dataset type and details on its location and schema.")
		cleanroomsml_createTrainingDatasetCmd.MarkFlagRequired("name")
		cleanroomsml_createTrainingDatasetCmd.MarkFlagRequired("role-arn")
		cleanroomsml_createTrainingDatasetCmd.MarkFlagRequired("training-data")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_createTrainingDatasetCmd)
}
