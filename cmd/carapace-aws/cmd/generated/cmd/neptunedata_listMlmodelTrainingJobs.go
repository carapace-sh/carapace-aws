package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_listMlmodelTrainingJobsCmd = &cobra.Command{
	Use:   "list-mlmodel-training-jobs",
	Short: "Lists Neptune ML model-training jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_listMlmodelTrainingJobsCmd).Standalone()

	neptunedata_listMlmodelTrainingJobsCmd.Flags().String("max-items", "", "The maximum number of items to return (from 1 to 1024; the default is 10).")
	neptunedata_listMlmodelTrainingJobsCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
	neptunedataCmd.AddCommand(neptunedata_listMlmodelTrainingJobsCmd)
}
