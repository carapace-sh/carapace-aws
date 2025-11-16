package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getMlmodelTransformJobCmd = &cobra.Command{
	Use:   "get-mlmodel-transform-job",
	Short: "Gets information about a specified model transform job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getMlmodelTransformJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_getMlmodelTransformJobCmd).Standalone()

		neptunedata_getMlmodelTransformJobCmd.Flags().String("id", "", "The unique identifier of the model-transform job to be reetrieved.")
		neptunedata_getMlmodelTransformJobCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
		neptunedata_getMlmodelTransformJobCmd.MarkFlagRequired("id")
	})
	neptunedataCmd.AddCommand(neptunedata_getMlmodelTransformJobCmd)
}
