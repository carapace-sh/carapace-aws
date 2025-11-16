package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_cancelMldataProcessingJobCmd = &cobra.Command{
	Use:   "cancel-mldata-processing-job",
	Short: "Cancels a Neptune ML data processing job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_cancelMldataProcessingJobCmd).Standalone()

	neptunedata_cancelMldataProcessingJobCmd.Flags().Bool("clean", false, "If set to `TRUE`, this flag specifies that all Neptune ML S3 artifacts should be deleted when the job is stopped.")
	neptunedata_cancelMldataProcessingJobCmd.Flags().String("id", "", "The unique identifier of the data-processing job.")
	neptunedata_cancelMldataProcessingJobCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
	neptunedata_cancelMldataProcessingJobCmd.Flags().Bool("no-clean", false, "If set to `TRUE`, this flag specifies that all Neptune ML S3 artifacts should be deleted when the job is stopped.")
	neptunedata_cancelMldataProcessingJobCmd.MarkFlagRequired("id")
	neptunedata_cancelMldataProcessingJobCmd.Flag("no-clean").Hidden = true
	neptunedataCmd.AddCommand(neptunedata_cancelMldataProcessingJobCmd)
}
