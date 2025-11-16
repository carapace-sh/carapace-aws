package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_listMlmodelTransformJobsCmd = &cobra.Command{
	Use:   "list-mlmodel-transform-jobs",
	Short: "Returns a list of model transform job IDs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_listMlmodelTransformJobsCmd).Standalone()

	neptunedata_listMlmodelTransformJobsCmd.Flags().String("max-items", "", "The maximum number of items to return (from 1 to 1024; the default is 10).")
	neptunedata_listMlmodelTransformJobsCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
	neptunedataCmd.AddCommand(neptunedata_listMlmodelTransformJobsCmd)
}
