package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_listMldataProcessingJobsCmd = &cobra.Command{
	Use:   "list-mldata-processing-jobs",
	Short: "Returns a list of Neptune ML data processing jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_listMldataProcessingJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedata_listMldataProcessingJobsCmd).Standalone()

		neptunedata_listMldataProcessingJobsCmd.Flags().String("max-items", "", "The maximum number of items to return (from 1 to 1024; the default is 10).")
		neptunedata_listMldataProcessingJobsCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
	})
	neptunedataCmd.AddCommand(neptunedata_listMldataProcessingJobsCmd)
}
