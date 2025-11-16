package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeDataDeletionJobCmd = &cobra.Command{
	Use:   "describe-data-deletion-job",
	Short: "Describes the data deletion job created by [CreateDataDeletionJob](https://docs.aws.amazon.com/personalize/latest/dg/API_CreateDataDeletionJob.html), including the job status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeDataDeletionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_describeDataDeletionJobCmd).Standalone()

		personalize_describeDataDeletionJobCmd.Flags().String("data-deletion-job-arn", "", "The Amazon Resource Name (ARN) of the data deletion job.")
		personalize_describeDataDeletionJobCmd.MarkFlagRequired("data-deletion-job-arn")
	})
	personalizeCmd.AddCommand(personalize_describeDataDeletionJobCmd)
}
