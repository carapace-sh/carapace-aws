package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_getMldataProcessingJobCmd = &cobra.Command{
	Use:   "get-mldata-processing-job",
	Short: "Retrieves information about a specified data processing job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_getMldataProcessingJobCmd).Standalone()

	neptunedata_getMldataProcessingJobCmd.Flags().String("id", "", "The unique identifier of the data-processing job to be retrieved.")
	neptunedata_getMldataProcessingJobCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
	neptunedata_getMldataProcessingJobCmd.MarkFlagRequired("id")
	neptunedataCmd.AddCommand(neptunedata_getMldataProcessingJobCmd)
}
