package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listLabelingJobsForWorkteamCmd = &cobra.Command{
	Use:   "list-labeling-jobs-for-workteam",
	Short: "Gets a list of labeling jobs assigned to a specified work team.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listLabelingJobsForWorkteamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listLabelingJobsForWorkteamCmd).Standalone()

		sagemaker_listLabelingJobsForWorkteamCmd.Flags().String("creation-time-after", "", "A filter that returns only labeling jobs created after the specified time (timestamp).")
		sagemaker_listLabelingJobsForWorkteamCmd.Flags().String("creation-time-before", "", "A filter that returns only labeling jobs created before the specified time (timestamp).")
		sagemaker_listLabelingJobsForWorkteamCmd.Flags().String("job-reference-code-contains", "", "A filter the limits jobs to only the ones whose job reference code contains the specified string.")
		sagemaker_listLabelingJobsForWorkteamCmd.Flags().String("max-results", "", "The maximum number of labeling jobs to return in each page of the response.")
		sagemaker_listLabelingJobsForWorkteamCmd.Flags().String("next-token", "", "If the result of the previous `ListLabelingJobsForWorkteam` request was truncated, the response includes a `NextToken`.")
		sagemaker_listLabelingJobsForWorkteamCmd.Flags().String("sort-by", "", "The field to sort results by.")
		sagemaker_listLabelingJobsForWorkteamCmd.Flags().String("sort-order", "", "The sort order for results.")
		sagemaker_listLabelingJobsForWorkteamCmd.Flags().String("workteam-arn", "", "The Amazon Resource Name (ARN) of the work team for which you want to see labeling jobs for.")
		sagemaker_listLabelingJobsForWorkteamCmd.MarkFlagRequired("workteam-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_listLabelingJobsForWorkteamCmd)
}
