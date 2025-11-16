package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getMlmodelTrainingJobCmd = &cobra.Command{
	Use:   "get-mlmodel-training-job",
	Short: "Retrieves information about a Neptune ML model training job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getMlmodelTrainingJobCmd).Standalone()

	neptunedata_getMlmodelTrainingJobCmd.Flags().String("id", "", "The unique identifier of the model-training job to retrieve.")
	neptunedata_getMlmodelTrainingJobCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
	neptunedata_getMlmodelTrainingJobCmd.MarkFlagRequired("id")
	neptunedataCmd.AddCommand(neptunedata_getMlmodelTrainingJobCmd)
}
