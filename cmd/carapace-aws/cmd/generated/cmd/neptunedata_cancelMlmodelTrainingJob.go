package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_cancelMlmodelTrainingJobCmd = &cobra.Command{
	Use:   "cancel-mlmodel-training-job",
	Short: "Cancels a Neptune ML model training job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_cancelMlmodelTrainingJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_cancelMlmodelTrainingJobCmd).Standalone()

		neptunedata_cancelMlmodelTrainingJobCmd.Flags().Bool("clean", false, "If set to `TRUE`, this flag specifies that all Amazon S3 artifacts should be deleted when the job is stopped.")
		neptunedata_cancelMlmodelTrainingJobCmd.Flags().String("id", "", "The unique identifier of the model-training job to be canceled.")
		neptunedata_cancelMlmodelTrainingJobCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
		neptunedata_cancelMlmodelTrainingJobCmd.Flags().Bool("no-clean", false, "If set to `TRUE`, this flag specifies that all Amazon S3 artifacts should be deleted when the job is stopped.")
		neptunedata_cancelMlmodelTrainingJobCmd.MarkFlagRequired("id")
		neptunedata_cancelMlmodelTrainingJobCmd.Flag("no-clean").Hidden = true
	})
	neptunedataCmd.AddCommand(neptunedata_cancelMlmodelTrainingJobCmd)
}
