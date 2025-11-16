package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_cancelMlmodelTransformJobCmd = &cobra.Command{
	Use:   "cancel-mlmodel-transform-job",
	Short: "Cancels a specified model transform job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_cancelMlmodelTransformJobCmd).Standalone()

	neptunedata_cancelMlmodelTransformJobCmd.Flags().Bool("clean", false, "If this flag is set to `TRUE`, all Neptune ML S3 artifacts should be deleted when the job is stopped.")
	neptunedata_cancelMlmodelTransformJobCmd.Flags().String("id", "", "The unique ID of the model transform job to be canceled.")
	neptunedata_cancelMlmodelTransformJobCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
	neptunedata_cancelMlmodelTransformJobCmd.Flags().Bool("no-clean", false, "If this flag is set to `TRUE`, all Neptune ML S3 artifacts should be deleted when the job is stopped.")
	neptunedata_cancelMlmodelTransformJobCmd.MarkFlagRequired("id")
	neptunedata_cancelMlmodelTransformJobCmd.Flag("no-clean").Hidden = true
	neptunedataCmd.AddCommand(neptunedata_cancelMlmodelTransformJobCmd)
}
